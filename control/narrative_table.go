package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/opencontrols"
)

func fillNarrativeRows(rows []xml.Node, data opencontrols.Data, control string) error {
	for _, row := range rows {
		section := narrativeSection{row}
		err := section.Fill(data, control)
		if err != nil {
			return err
		}
	}
	return nil
}

// NarrativeTable represents the node in the Word docx XML tree that corresponds to the justification fields for a security control.
type NarrativeTable struct {
	table
}

// NewNarrativeTable creates a NarrativeTable instance.
func NewNarrativeTable(root xml.Node) NarrativeTable {
	tbl := table{Root: root}
	return NarrativeTable{tbl}
}

// SectionRows returns the list of rows which correspond to each "part" of the narrative. Will return a single row when the narrative isn't split into parts.
func (t *NarrativeTable) SectionRows() ([]xml.Node, error) {
	// skip the header row
	return t.table.searchSubtree(`.//w:tr[position() > 1]`)
}

// Fill inserts the OpenControl data into the table.
func (t *NarrativeTable) Fill(openControlData opencontrols.Data) (err error) {
	control, err := t.table.controlName()
	if err != nil {
		return
	}

	rows, err := t.SectionRows()
	if err != nil {
		return
	}

	fillNarrativeRows(rows, openControlData, control)
	return
}
