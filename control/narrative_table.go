package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/opencontrols"
)

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
		textFields, err := row.Search(`(./w:tc/w:p)[1]`)
		if err != nil {
			return err
		}
		textField := textFields[0]

		narrative := openControlData.GetNarrative(control, "")
		textField.SetContent(narrative)
	} else {
		// multiple parts
		for _, row := range rows {
			// TODO remove hard-coding
			sectionKey := "b"

			textFields, err := row.Search(`./w:tc[position() = 1]/w:p`)
			if err != nil {
				return err
			}
			textField := textFields[0]

			narrative := openControlData.GetNarrative(control, sectionKey)
			textField.SetContent(narrative)
		}
	}

	return
}
