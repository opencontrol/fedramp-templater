package opencontrols

import (
	"github.com/opencontrol/compliance-masonry/commands/docs/docx"
	"github.com/opencontrol/compliance-masonry/models"
)

// Data contains the OpenControl justification information.
type Data struct {
	ocd docx.OpenControlDocx
}

// LoadFrom creates a new Data struct from the provided path to an `opencontrols/` directory.
func LoadFrom(path string) (data Data, errors []error) {
	openControlData, errors := models.LoadData(path, "")
	if len(errors) > 0 {
		return
	}

	ocd := docx.OpenControlDocx{openControlData}
	data = Data{ocd}
	return
}

// GetResponsibleRoles returns the responsible role information for each component matching the specified control.
func (d *Data) GetResponsibleRoles(control string) string {
	return d.ocd.FormatResponsibleRoles("NIST-800-53", control)
}
