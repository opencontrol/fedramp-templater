package opencontrols

import (
	"github.com/opencontrol/compliance-masonry/commands/docs/docx"
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/fedramp-templater/common/origin"
	"github.com/opencontrol/fedramp-templater/common/status"
	"gopkg.in/fatih/set.v0"
	"strings"
)

const standardKey = "NIST-800-53"

// Data contains the OpenControl justification information.
type Data struct {
	ocd docx.OpenControlDocx
}

// LoadFrom creates a new Data struct from the provided path to an `opencontrols/` directory.
func LoadFrom(dirPath string) (data Data, errors []error) {
	openControlData, errors := models.LoadData(dirPath, "")
	if len(errors) > 0 {
		return
	}

	ocd := docx.OpenControlDocx{openControlData}
	data = Data{ocd}
	return
}

// GetResponsibleRoles returns the responsible role information for each component matching the specified control.
func (d *Data) GetResponsibleRoles(control string) string {
	var responsibleRoleOrig = d.ocd.FormatResponsibleRoles(standardKey, control)
	if len(strings.TrimSpace(responsibleRoleOrig)) > 0 {
		responsibleRoleOrig += "\n"
	}
	return responsibleRoleOrig
}

// mergeNewLines - replace single newlines with space to preserve Word line layout
func mergeNewLines(text string) string {
	// initialize result and other vars
	var (
		result      = ""
		isFirstLine = true
		hasPrevLine = false
	)

	// split text into lines and process
	lines := strings.Split(text, "\n")
	for i := range lines {
		line := lines[i]

		// case 1: for first line, accept as-is
		if isFirstLine {
			result = line + "\n"
			isFirstLine = false
			continue
		}

		// case 2: line is empty (indicates start new pp)
		if len(strings.TrimSpace(line)) == 0 {
			result += "\n"
			hasPrevLine = false
			continue
		}

		// case 3: append space if previous line (not newline)
		if hasPrevLine {
			result += " "
		}
		result += line

		// permit lines to be continued
		hasPrevLine = true
	}

	// account for trailing last line; auto-append newline
	if hasPrevLine {
		result += "\n"
	}
	return result
}

// GetNarrative returns the justification text for the specified control. Pass an empty string for `sectionKey` if you are looking for the overall narrative.
func (d *Data) GetNarrative(control string, sectionKey string) string {
	var narrative = d.ocd.FormatNarrative(standardKey, control, sectionKey)
	return mergeNewLines(narrative)
}

// GetParameter returns the justification text for the specified control. Pass an empty string for `sectionKey` if you are looking for the overall narrative.
func (d *Data) GetParameter(control string, sectionKey string) string {

	var parameter = d.ocd.FormatParameter(standardKey, control, sectionKey)
	return mergeNewLines(parameter)
}

// GetControlOrigins returns the control origination information for each component matching the specified control.
func (d *Data) GetControlOrigins(control string) ControlOrigins {
	controlOrigins := ControlOrigins{}
	justifications := d.ocd.Justifications.Get(standardKey, control)
	for _, justification := range justifications {
		numberOfControlOrigins := len(justification.SatisfiesData.GetControlOrigins())
		if numberOfControlOrigins > 1 {
			for _, orgs := range justification.SatisfiesData.GetControlOrigins() {
				controlOrigins.origins = append(controlOrigins.origins, orgs)
			}

		} else {
			numberOfControlOrigin := len(justification.SatisfiesData.GetControlOrigin())
			if numberOfControlOrigin != 0 {
				controlOrigins.origins = append(controlOrigins.origins, justification.SatisfiesData.GetControlOrigin())
			} else {
				controlOrigins.origins = append(controlOrigins.origins, justification.SatisfiesData.GetControlOrigins()[0])
			}

		}
	}

	return controlOrigins
}

// ControlOrigins is a wrapper for the extracted data from the YAML for a particular control.
type ControlOrigins struct {
	origins []string
}

func detectControlOriginKey(text string) origin.Key {
	controlOriginMappings := origin.GetSourceMappings()
	for controlOrigin, controlOriginMapping := range controlOriginMappings {
		if controlOriginMapping.IsYAMLMappingEqualTo(text) {
			return controlOrigin
		}
	}
	return origin.NoOrigin
}

// GetCheckedOrigins will return the list of origin keys.
func (origins ControlOrigins) GetCheckedOrigins() *set.Set {
	// find the control origins currently checked in the section in the YAML.
	yamlControlOrigins := set.New()
	for _, controlOrigin := range origins.origins {
		controlOriginKey := detectControlOriginKey(controlOrigin)
		if controlOriginKey == origin.NoOrigin {
			continue
		}
		yamlControlOrigins.Add(controlOriginKey)

	}
	return yamlControlOrigins
}

// GetImplementationStatuses returns the control origination information for each component matching the specified control.
func (d *Data) GetImplementationStatuses(control string) ImplementationStatuses {
	implementationStatuses := ImplementationStatuses{}
	justifications := d.ocd.Justifications.Get(standardKey, control)
	for _, justification := range justifications {
		numberOfImplementationStatuses := len(justification.SatisfiesData.GetImplementationStatuses())
		if numberOfImplementationStatuses > 1 {
			for _, stats := range justification.SatisfiesData.GetImplementationStatuses() {
				implementationStatuses.statuses = append(implementationStatuses.statuses, stats)
			}

		} else {
			numberOfImplementationStatus := len(justification.SatisfiesData.GetImplementationStatus())
			if numberOfImplementationStatus != 0 {
				implementationStatuses.statuses = append(implementationStatuses.statuses, justification.SatisfiesData.GetImplementationStatus())
			} else {
				implementationStatuses.statuses = append(implementationStatuses.statuses, justification.SatisfiesData.GetImplementationStatuses()[0])
			}

		}
	}
	return implementationStatuses
}

// ImplementationStatuses is a wrapper for the extracted data from the YAML for a particular control.
type ImplementationStatuses struct {
	statuses []string
}

func detectImplementationStatusKey(text string) status.Key {
	implementationStatusMappings := status.GetSourceMappings()
	for implementationStatus, implementationStatusMapping := range implementationStatusMappings {
		if implementationStatusMapping.IsYAMLMappingEqualTo(text) {
			return implementationStatus
		}
	}
	return status.NoStatus
}

//GetCheckedStatuses will return the list of status keys.
func (statuses ImplementationStatuses) GetCheckedStatuses() *set.Set {
	// find the implementation statuses currently checked in the section in the YAML.
	yamlImplementationStatuses := set.New()
	for _, implementationStatus := range statuses.statuses {
		implementationStatusKey := detectImplementationStatusKey(implementationStatus)
		if implementationStatusKey == status.NoStatus {
			continue
		}
		yamlImplementationStatuses.Add(implementationStatusKey)

	}
	return yamlImplementationStatuses
}
