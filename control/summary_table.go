package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/common/origin"
	"github.com/opencontrol/fedramp-templater/common/source"
	"github.com/opencontrol/fedramp-templater/common/status"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/reporter"
	"gopkg.in/fatih/set.v0"
	"fmt"
)

const (
	responsibleRoleField      = "Responsible Role"
	controlOriginationField   = "Control Origination"
	implementationStatusField = "Implementation Status"
)

// SummaryTable represents the node in the Word docx XML tree that corresponds to the summary information for a security control.
type SummaryTable struct {
	table
	originTable *controlOrigination
	statusTable *implementationStatus
}

// NewSummaryTable creates a SummaryTable instance.
func NewSummaryTable(root xml.Node) (SummaryTable, error) {
	tbl := table{Root: root}
	originTable, err := newControlOrigination(&tbl)
	if err != nil {
		return SummaryTable{}, err
	}
	statusTable, err := newImplementationStatus(&tbl)
	if err != nil {
		return SummaryTable{}, err
	}
	return SummaryTable{tbl, originTable, statusTable}, nil
}

func (st *SummaryTable) controlName() (name string, err error) {
	return st.table.controlName()
}

// ControlName - name of the control
func (st *SummaryTable) ControlName() (name string, err error) {
	return st.table.controlName()
}

func (st *SummaryTable) fillResponsibleRole(openControlData opencontrols.Data, control string) (err error) {
	roleCell, err := findResponsibleRole(st)
	if err != nil {
		return
	}

	roles := openControlData.GetResponsibleRoles(control)
	roleCell.setValue(roles)
	return
}

func (st *SummaryTable) fillControlOrigination(openControlData opencontrols.Data, control string) (err error) {
	controlOrigins := openControlData.GetControlOrigins(control)
	checkedOriginsSet := controlOrigins.GetCheckedOrigins()
	checkedOrigins := origin.ConvertSetToKeys(checkedOriginsSet)

	for _, checkedOrigin := range checkedOrigins {
		if checkedOrigin == origin.NoOrigin {
			continue
		}
		st.originTable.origins[checkedOrigin].SetCheckMarkTo(true)
	}
	return
}

func (st *SummaryTable) fillImplementationStatus(openControlData opencontrols.Data, control string) (err error) {
	implementationStatuses := openControlData.GetImplementationStatuses(control)
	checkedStatusesSet := implementationStatuses.GetCheckedStatuses()
	checkedStatuses := status.ConvertSetToKeys(checkedStatusesSet)

	for _, checkedStatus := range checkedStatuses {
		if checkedStatus == status.NoStatus {
			continue
		}
		st.statusTable.statuses[checkedStatus].SetCheckMarkTo(true)
	}
	return
}

// Fill inserts the OpenControl justifications into the table. Note this modifies the `table`.
func (st *SummaryTable) Fill(openControlData opencontrols.Data) (err error) {
	control, err := st.controlName()
	if err != nil {
		return
	}
	err = st.fillResponsibleRole(openControlData, control)
	if err != nil {
		return
	}
	fmt.Println("Filling controls for ", control)
	err = st.fillControlOrigination(openControlData, control)
	if err != nil {
		return
	}
	err = st.fillImplementationStatus(openControlData, control)
	if err != nil {
		return
	}

	return
}

// diffControlOrigination computes the diff of the control origination.
func (st *SummaryTable) diffControlOrigination(control string,
	openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	// find the control origins currently checked in the section in the doc.
	docControlOrigins := st.originTable.getCheckedOrigins()

	// find the control origins noted in the yaml.
	yamlControlOriginationData := openControlData.GetControlOrigins(control)
	// find the control origins currently checked in the section in the YAML.
	yamlControlOrigins := yamlControlOriginationData.GetCheckedOrigins()

	// find the difference of the two sets.
	controlOriginMap := origin.GetSourceMappings()
	reports := []reporter.Reporter{}

	// find only the origins in the document.
	onlyInDocOrigins := set.Difference(docControlOrigins, yamlControlOrigins)
	// create the diff report for the origins only in the document.
	onlyInDocOriginReports := st.createControlOriginsDiffReport(onlyInDocOrigins, controlOriginMap, control, source.SSP)
	reports = append(reports, onlyInDocOriginReports...)

	// find only the origins in the yaml.
	onlyInYAMLOrigins := set.Difference(yamlControlOrigins, docControlOrigins)
	// create the diff report for the origins only in the yaml.
	onlyInYAMLOriginReports := st.createControlOriginsDiffReport(onlyInYAMLOrigins, controlOriginMap, control, source.YAML)
	reports = append(reports, onlyInYAMLOriginReports...)

	return reports, nil
}

// diffImplementationStatus computes the diff of the implementation status.
func (st *SummaryTable) diffImplementationStatus(control string,
	openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	// find the implementation statues currently checked in the section in the doc.
	docImplementationStatuses := st.statusTable.getCheckedStatuses()

	// find the implementation statuses noted in the yaml.
	yamlImplementationStatusData := openControlData.GetImplementationStatuses(control)
	// find the implementation statuses currently checked in the section in the YAML.
	yamlImplementationStatuses := yamlImplementationStatusData.GetCheckedStatuses()

	// find the difference of the two sets.
	implementationStatusMap := status.GetSourceMappings()
	reports := []reporter.Reporter{}

	// find only the statuses in the document.
	onlyInDocStatuses := set.Difference(docImplementationStatuses, yamlImplementationStatuses)
	// create the diff report for the statuses only in the document.
	onlyInDocStatusReports := st.createImplementationStatusesDiffReport(onlyInDocStatuses, implementationStatusMap, control, source.SSP)
	reports = append(reports, onlyInDocStatusReports...)

	// find only the origins in the yaml.
	onlyInYAMLStatuses := set.Difference(yamlImplementationStatuses, docImplementationStatuses)
	// create the diff report for the origins only in the yaml.
	onlyInYAMLStatusReports := st.createImplementationStatusesDiffReport(onlyInYAMLStatuses, implementationStatusMap, control, source.YAML)
	reports = append(reports, onlyInYAMLStatusReports...)

	return reports, nil
}

func (*SummaryTable) createControlOriginsDiffReport(diff set.Interface,
	controlOriginSrcMap map[origin.Key]origin.SrcMapping, control string, src source.Source) []reporter.Reporter {
	reports := []reporter.Reporter{}
	secondField := field{text: ""}
	originKeys := origin.ConvertSetToKeys(diff)
	for _, originKey := range originKeys {
		var firstField field
		switch src {
		case source.SSP:
			firstField.text = controlOriginSrcMap[originKey][source.SSP]
			firstField.source = source.SSP
			secondField.source = source.YAML
		case source.YAML:
			firstField.text = controlOriginSrcMap[originKey][source.YAML]
			firstField.source = source.YAML
			secondField.source = source.SSP
		}
		// Get the doc mapping and put it in the doc.
		reports = append(reports, NewDiff(control, controlOriginationField, firstField, secondField))
	}
	return reports
}

func (*SummaryTable) createImplementationStatusesDiffReport(diff set.Interface,
	implementationStatusSrcMap map[status.Key]status.SrcMapping, control string, src source.Source) []reporter.Reporter {
	reports := []reporter.Reporter{}
	secondField := field{text: ""}
	statusKeys := status.ConvertSetToKeys(diff)
	for _, statusKey := range statusKeys {
		var firstField field
		switch src {
		case source.SSP:
			firstField.text = implementationStatusSrcMap[statusKey][source.SSP]
			firstField.source = source.SSP
			secondField.source = source.YAML
		case source.YAML:
			firstField.text = implementationStatusSrcMap[statusKey][source.YAML]
			firstField.source = source.YAML
			secondField.source = source.SSP
		}
		// Get the doc mapping and put it in the doc.
		reports = append(reports, NewDiff(control, implementationStatusField, firstField, secondField))
	}
	return reports
}

// diffResponsibleRole computes the diff of the responsible role cell.
func (st *SummaryTable) diffResponsibleRole(control string, openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	roleCell, err := findResponsibleRole(st)
	if err != nil {
		return []reporter.Reporter{}, err
	}
	yamlField := field{source: source.YAML}
	yamlField.text = openControlData.GetResponsibleRoles(control)
	sspField := field{source: source.SSP}
	sspField.text = roleCell.getValue()
	if roleCell.isDefaultValue(sspField.text) || yamlField.text == sspField.text {
		return []reporter.Reporter{}, nil
	}
	return []reporter.Reporter{
		NewDiff(control, responsibleRoleField, sspField, yamlField),
	}, nil
}

// Diff returns the list of diffs in the control table.
func (st *SummaryTable) Diff(openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	reports := []reporter.Reporter{}
	control, err := st.controlName()
	if err != nil {
		return reports, err
	}
	// Diff the responsible roles
	diffReports, err := st.diffResponsibleRole(control, openControlData)
	if err != nil {
		return reports, err
	}
	reports = append(reports, diffReports...)

	// Diff the control origination
	diffReports, err = st.diffControlOrigination(control, openControlData)
	if err != nil {
		return reports, err
	}
	reports = append(reports, diffReports...)
	// Diff the implememtation status
	diffReports, err = st.diffImplementationStatus(control, openControlData)
	if err != nil {
		return reports, err
	}
	reports = append(reports, diffReports...)
	return reports, nil
}
