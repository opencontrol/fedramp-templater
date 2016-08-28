package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/opencontrols"
)

func fillRow(row xml.Node, data opencontrols.Data, control string, section string) {
	// equivalent XPath: `./w:tc[last()]/w:p[1]`
	textField := row.LastChild().FirstChild()
	narrative := data.GetNarrative(control, section)
	textField.SetContent(narrative)
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
		for _, row := range rows {
			// TODO remove hard-coding
			sectionKey := "b"
			fillRow(row, openControlData, control, sectionKey)
		}
	}

	return
}
