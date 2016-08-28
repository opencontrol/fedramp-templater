package control

import (
	"errors"
	"regexp"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/docx/helper"
	"github.com/opencontrol/fedramp-templater/opencontrols"
)

func findSectionKey(row xml.Node) (section string, err error) {
	re := regexp.MustCompile(`Part ([a-z])`)
	subMatches := re.FindSubmatch([]byte(row.Content()))
	if len(subMatches) != 2 {
		err = errors.New("No Parts found.")
		return
	}
	section = string(subMatches[1])
	return
}

func fillRow(row xml.Node, data opencontrols.Data, control string, section string) (err error) {
	// the row should have one or two cells; either way, the last one is what should be filled
	paragraphNodes, err := row.Search(`./w:tc[last()]/w:p[1]`)
	if err != nil {
		return
	}
	paragraphNode := paragraphNodes[0]

	narrative := data.GetNarrative(control, section)
	helper.FillParagraph(paragraphNode, narrative)

	return
}

func fillRows(rows []xml.Node, data opencontrols.Data, control string) error {
	for _, row := range rows {
		sectionKey, err := findSectionKey(row)
		if err != nil {
			return err
		}

		err = fillRow(row, data, control, sectionKey)
		if err != nil {
			return err
		}
	}
	return nil
}

type NarrativeTable struct {
	tbl table
}

func NewNarrativeTable(root xml.Node) NarrativeTable {
	tbl := table{Root: root}
	return NarrativeTable{tbl}
}

func (t *NarrativeTable) SectionRows() ([]xml.Node, error) {
	// skip the header row
	return t.tbl.searchSubtree(`.//w:tr[position() > 1]`)
}

func (t *NarrativeTable) Fill(openControlData opencontrols.Data) (err error) {
	control, err := t.tbl.controlName()
	if err != nil {
		return
	}

	rows, err := t.SectionRows()
	if err != nil {
		return
	}

	if len(rows) == 1 {
		// singular narrative
		row := rows[0]
		fillRow(row, openControlData, control, "")
	} else {
		// multiple parts
		fillRows(rows, openControlData, control)
	}

	return
}
