package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/reporter"
	"gopkg.in/fatih/set.v0"
	"log"
	"reflect"
)

const (
	responsibleRoleField    = "Responsible Role"
	controlOriginationField = "Control Origination"
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

func (st *SummaryTable) controlName() (name string, err error) {
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
	controlOrigination, err := newControlOrigination(st)
	if err != nil {
		return
	}

	controlOrigins := openControlData.GetControlOrigins(control)
	for _, controlOrigin := range controlOrigins {
		controlOriginKey := detectControlOriginKeyFromYAML(controlOrigin)
		if controlOriginKey == noOrigin {
			continue
		}
		controlOrigination.origins[controlOriginKey].SetCheckMarkTo(true)
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
	err = st.fillControlOrigination(openControlData, control)
	if err != nil {
		return
	}

	return
}

// diffControlOrigination computes the diff of the control origination.
func (st *SummaryTable) diffControlOrigination(control string,
	openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	// find the control origination section in the document.
	docControlOriginationData, err := newControlOrigination(st)
	if err != nil {
		return nil, err
	}
	// find the control origins currently checked in the section in the doc.
	docControlOrigins := docControlOriginationData.getCheckedOrigins()

	// find the control origins noted in the yaml.
	yamlControlOriginationData := openControlData.GetControlOrigins(control)
	// find the control origins currently checked in the section in the YAML.
	yamlControlOrigins := getCheckedOriginsFromYAML(yamlControlOriginationData)

	// find the difference of the two sets.
	controlOriginMap := getControlOriginMappings()
	reports := []reporter.Reporter{}

	// find only the origins in the document.
	onlyInDocOrigins := set.Difference(docControlOrigins, yamlControlOrigins)
	// create the diff report for the origins only in the document.
	onlyInDocOriginReports := createControlOriginsDiffReport(onlyInDocOrigins, controlOriginMap, control, sspSrc)
	reports = append(reports, onlyInDocOriginReports...)

	// find only the origins in the yaml.
	onlyInYAMLOrigins := set.Difference(yamlControlOrigins, docControlOrigins)
	// create the diff report for the origins only in the yaml.
	onlyInYAMLOriginReports := createControlOriginsDiffReport(onlyInYAMLOrigins, controlOriginMap, control, yamlSrc)
	reports = append(reports, onlyInYAMLOriginReports...)

	return reports, nil
}

func createControlOriginsDiffReport(diff set.Interface, controlOriginMap map[controlOrigin]originMapping,
	control string, source infoSource) []reporter.Reporter {
	reports := []reporter.Reporter{}
	secondField := field{text: ""}
	for _, originInterface := range diff.List() {
		// cast back from interface{} to controlOrigin so we can use it in the controlOriginMap
		origin, isType := originInterface.(controlOrigin)
		if isType {
			var firstField field
			switch source {
			case sspSrc:
				firstField.text = controlOriginMap[origin][sspSrc]
				firstField.source = sspSrc
				secondField.source = yamlSrc
			case yamlSrc:
				firstField.text = controlOriginMap[origin][yamlSrc]
				firstField.source = yamlSrc
				secondField.source = sspSrc
			}
			// Get the doc mapping and put it in the doc.
			reports = append(reports, NewDiff(control, controlOriginationField, firstField, secondField))
		} else {
			log.Printf("Unable to use value as 'controlOrigin' Type: %v. Value: %v.\n",
				reflect.TypeOf(originInterface), originInterface)
		}
	}
	return reports
}

// diffResponsibleRole computes the diff of the responsible role cell.
func (st *SummaryTable) diffResponsibleRole(control string, openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	roleCell, err := findResponsibleRole(st)
	if err != nil {
		return []reporter.Reporter{}, err
	}
	yamlField := field{source: yamlSrc}
	yamlField.text = openControlData.GetResponsibleRoles(control)
	sspField := field{source: sspSrc}
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
	return reports, nil
}
