package control

import (
	"fmt"
	"github.com/opencontrol/fedramp-templater/reporter"
	"io"
	"strings"
)

type diffReporter struct {
	controlName string
	fieldType   string
	firstField  field
	secondField field
}

// NewDiff creates a new collection of information that can report diff info for a control.
func NewDiff(controlName, fieldType string, firstField, secondField field) reporter.Reporter {
	return diffReporter{
		controlName: strings.TrimSpace(controlName),
		fieldType:   strings.TrimSpace(fieldType),
		firstField:  firstField,
		secondField: secondField,
	}
}

// WriteTextTo writes diff information for a control to the writer in plain text format.
func (r diffReporter) WriteTextTo(writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "Control: %s. %s in %s: \"%s\". %s in %s: \"%s\".\n",
		r.controlName, r.fieldType, r.firstField.source, strings.TrimSpace(r.firstField.text),
		r.fieldType, r.secondField.source, strings.TrimSpace(r.secondField.text))
	return err
}
