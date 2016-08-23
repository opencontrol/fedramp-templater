package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"errors"
	"regexp"
	"strings"
)

// FindResponsibleRole looks for the Responsible Role cell in the control table.
func FindResponsibleRole(ct *Table) (*ResponsibleRole, error) {
	nodes, err := ct.searchSubtree(".//w:tc//w:t[contains(., 'Responsible Role')]")
	if err != nil {
		return nil, err
	}
	if len(nodes) != 1 {
		return nil, errors.New("could not find Responsible Role cell")
	}
	return &ResponsibleRole{node: nodes[0]}, nil
}

// ResponsibleRole is the container for the responsible role cell.
type ResponsibleRole struct {
	node xml.Node
}

// GetContent returns the full string representation of the content of the cell itself.
func (r *ResponsibleRole) GetContent() string {
	return r.node.Content()
}

// SetContent overrides the string representation of the content cell itself.
func (r *ResponsibleRole) SetContent(content string) {
	r.node.SetContent(content)
}

// IsDefaultValue contains the logic to detect if the input is a default value. This is looking at the extracted
// value of responsible role and not the full string representation.
func (r *ResponsibleRole) IsDefaultValue(value string) bool {
	return value == ""
}

// GetValue extracts the unique value from the full string representation.
// It looks at all the text after "Responsible Role:".
func (r *ResponsibleRole) GetValue() string {
	re := regexp.MustCompile("Responsible Role:?")
	// Get all the substrings
	subStrings := re.Split(r.node.Content(), -1)
	result := ""
	// Go through the substrings and find the first one that is not empty.
	// (So far has always been the string at index 1)
	for _, subString := range subStrings {
		// If we find an non-empty string, return it.
		trimmedString := strings.TrimSpace(subString)
		if len(trimmedString) > 0 {
			result = trimmedString
			break
		}
	}
	return result
}