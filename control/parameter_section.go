package control

import (
	"errors"
	"fmt"
	"github.com/jbowtie/gokogiri/xml"
	"regexp"
	"strings"

	"github.com/opencontrol/fedramp-templater/opencontrols"
	xmlHelper "github.com/opencontrol/fedramp-templater/xml/helper"
)

type parameterSection struct {
	row xml.Node
}

func (n parameterSection) parsePart() (key string, err error) {
	re := regexp.MustCompile(`Parameter [A-Z]{2}|\([a-z]\)|\(\d{1,3}\)|(\-\d)|\d:`)
	content := []byte(n.row.Content())
	subMatches := re.FindSubmatch(content)
	if len(subMatches) == 0 {
		err = errors.New("no parts found")
		return
	}
	contentLength := len(content)
	contentEnd := content[10:contentLength]

	key = string(contentEnd)
	key = strings.Replace(key, ":", "", -1)
	key = strings.TrimSpace(key)
	return
}

// GetKey returns the parameter section key. `key` will be an empty string if there is no "Parameter".
func (n parameterSection) GetKey() (key string, err error) {
	cells, err := xmlHelper.SearchSubtree(n.row, `.//w:t`)
	numCells := len(cells)
	if numCells >= 1 {
		key, err = n.parsePart()
		if err != nil {
			return
		}
	} else {
		err = errors.New("don't know how to parse row")
	}

	return
}

// Fill populates the section/part with the description for this control part from the provided data.
func (n parameterSection) Fill(data opencontrols.Data, control string) (err error) {
	// the row should have one or two cells; either way, the last one is what should be filled

	cellNode, err := xmlHelper.SearchLast(n.row, `.//w:t[last()]`)
	if err != nil {
		return
	}

	key, err := n.GetKey()
	if err != nil {
		return
	}

	// fixup the parameter
	parameter := data.GetParameter(control, key)

	if strings.Contains(cellNode.String(), "Parameter") {
		cellNode.SetContent(fmt.Sprintf("Parameter %s: %s", key, parameter))

	} else {
		if strings.Contains(cellNode.String(), ">:<") {
			cellNode.SetContent(fmt.Sprintf(" : %s", parameter))
		} else {
			cellNode.SetContent(fmt.Sprintf(" %s", parameter))
		}

	}

	return
}
