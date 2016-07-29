package opencontrols

import (
	"github.com/opencontrol/compliance-masonry/commands/docs/docx"
	"github.com/opencontrol/compliance-masonry/models"
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
	return d.ocd.FormatResponsibleRoles(standardKey, control)
}

// GetNarrative returns the justification text for the specified control. Pass an empty string for `sectionKey` if you are looking for the overall narrative.
func (d *Data) GetNarrative(control string, sectionKey string) string {
	return d.ocd.FormatNarrative(standardKey, control, sectionKey)
}
