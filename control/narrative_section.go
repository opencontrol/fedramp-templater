package control

import (
	"errors"
	"github.com/jbowtie/gokogiri/xml"
	docxHelper "github.com/opencontrol/fedramp-templater/docx/helper"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	xmlHelper "github.com/opencontrol/fedramp-templater/xml/helper"
	"regexp"
)

type narrativeSection struct {
	row xml.Node
}

func (n narrativeSection) parsePart() (key string, err error) {
	re := regexp.MustCompile(`Part ([a-z])`)
	content := []byte(n.row.Content())
	subMatches := re.FindSubmatch(content)
	if len(subMatches) != 2 {
		err = errors.New("no parts found")
		return
	}
	key = string(subMatches[1])
	return
}

// GetKey returns the narrative section "part"/key. `key` will be an empty string if there is no "Part".
func (n narrativeSection) GetKey() (key string, err error) {
	cells, err := xmlHelper.SearchSubtree(n.row, `./w:tc`)
	numCells := len(cells)
	if numCells == 1 {
		// there is only a single narrative section
		key = ""
	} else if numCells == 2 {
		key, err = n.parsePart()
		if err != nil {
			return
		}
	} else {
		err = errors.New("don't know how to parse row")
	}

	return
}

// Fill populates the section/part with the narrative for this control part from the provided data.
func (n narrativeSection) Fill(data opencontrols.Data, control string) (err error) {
	// the row should have one or two cells; either way, the last one is what should be filled
	cellNode, err := xmlHelper.SearchOne(n.row, `./w:tc[last()]`)
	if err != nil {
		return
	}

	key, err := n.GetKey()
	if err != nil {
		return
	}

	// fixup the narrative
	narrative := data.GetNarrative(control, key)
	docxHelper.FillCell(cellNode, narrative)

	return
}
