package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/opencontrols"
)

func fillParameterRows(rows []xml.Node, data opencontrols.Data, control string) error {
	for _, row := range rows {
		section := parameterSection{row}
		err := section.Fill(data, control)
		if err != nil {
			return err
		}

	}
	return nil
}

// ParameterTable represents the node in the Word docx XML tree that corresponds to the justification fields for a security control.
type ParameterTable struct {
	table
}

// NewParameterTable creates a ParameterTable instance.
func NewParameterTable(root xml.Node) ParameterTable {
	tbl := table{Root: root}
	return ParameterTable{tbl}
}

// SectionRows returns the list of rows which correspond to each row containing the parameter. Will return a single row when the parameter isn't split into parts.
func (t *ParameterTable) SectionRows() ([]xml.Node, error) {
	// skip the header row that contains the control name
	return t.table.searchSubtree(".//w:tc[starts-with(normalize-space(.), 'Parameter')]")
}

// Fill inserts the OpenControl data into the table.
func (t *ParameterTable) Fill(openControlData opencontrols.Data) (err error) {
	control, err := t.table.controlName()
	if err != nil {
		return
	}

	rows, err := t.SectionRows()
	if err != nil {
		return
	}

	fillParameterRows(rows, openControlData, control)
	return
}
