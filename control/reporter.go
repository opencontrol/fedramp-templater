package control

import (
	"github.com/opencontrol/fedramp-templater/reporter"
	"io"
	"fmt"
	"strings"
)

type diffReporter struct {
	controlName string
	field string
	sspValue string
	yamlValue string
}

// NewDiff creates a new collection of information that can report diff info for a control.
func NewDiff(controlName, field, sspValue, yamlValue string) reporter.Reporter {
	return diffReporter{
		controlName: strings.TrimSpace(controlName),
		field: strings.TrimSpace(field),
		sspValue: strings.TrimSpace(sspValue),
		yamlValue: strings.TrimSpace(yamlValue),
	}
}

// WriteTextTo writes diff information for a control to the writer in plain text format.
func (r diffReporter) WriteTextTo(writer io.Writer) error {
	_ , err := fmt.Fprintf(writer, "Control: %s. %s in SSP: \"%s\". %s in YAML: \"%s\".\n",
		r.controlName, r.field, r.sspValue, r.field, r.yamlValue)
	return err
}