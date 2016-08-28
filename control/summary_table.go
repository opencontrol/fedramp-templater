package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/reporter"
)

const (
	responsibleRoleField = "Responsible Role"
)

// SummaryTable represents the node in the Word docx XML tree that corresponds to the summary information for a security control.
type SummaryTable struct {
	tbl table
}

// NewSummaryTable creates a SummaryTable instance.
func NewSummaryTable(root xml.Node) SummaryTable {
	tbl := table{Root: root}
	return SummaryTable{tbl}
}

func (st *SummaryTable) controlName() (name string, err error) {
	return st.tbl.controlName()
}

// Fill inserts the OpenControl justifications into the table. Note this modifies the `table`.
func (st *SummaryTable) Fill(openControlData opencontrols.Data) (err error) {
	roleCell, err := findResponsibleRole(st)
	if err != nil {
		return
	}

	control, err := st.controlName()
	if err != nil {
		return
	}

	roles := openControlData.GetResponsibleRoles(control)
	roleCell.setValue(roles)

	return
}

// diffResponsibleRole computes the diff of the responsible role cell.
func (st *SummaryTable) diffResponsibleRole(control string, openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	roleCell, err := findResponsibleRole(st)
	if err != nil {
		return []reporter.Reporter{}, err
	}
	yamlRoles := openControlData.GetResponsibleRoles(control)
	sspRoles := roleCell.getValue()
	if roleCell.isDefaultValue(sspRoles) || yamlRoles == sspRoles {
		return []reporter.Reporter{}, nil
	}
	return []reporter.Reporter{
		NewDiff(control, responsibleRoleField, sspRoles, yamlRoles),
	}, nil
}

// Diff returns the list of diffs in the control table.
func (st *SummaryTable) Diff(openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	control, err := st.controlName()
	if err != nil {
		return []reporter.Reporter{}, err
	}
	return st.diffResponsibleRole(control, openControlData)
}
