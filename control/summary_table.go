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
	table
}

// NewSummaryTable creates a SummaryTable instance.
func NewSummaryTable(root xml.Node) SummaryTable {
	tbl := table{Root: root}
	return SummaryTable{tbl}
}

func (ct *SummaryTable) controlName() (name string, err error) {
	return ct.table.controlName()
}

// Fill inserts the OpenControl justifications into the table. Note this modifies the `table`.
func (ct *SummaryTable) Fill(openControlData opencontrols.Data) (err error) {
	roleCell, err := findResponsibleRole(ct)
	if err != nil {
		return
	}

	control, err := ct.controlName()
	if err != nil {
		return
	}

	roles := openControlData.GetResponsibleRoles(control)
	roleCell.setValue(roles)

	return
}

// diffResponsibleRole computes the diff of the responsible role cell.
func (ct *SummaryTable) diffResponsibleRole(control string, openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	roleCell, err := findResponsibleRole(ct)
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
func (ct *SummaryTable) Diff(openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	control, err := ct.controlName()
	if err != nil {
		return []reporter.Reporter{}, err
	}
	return ct.diffResponsibleRole(control, openControlData)
}
