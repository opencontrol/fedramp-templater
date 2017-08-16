package control

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/xml/helper"
)

// findResponsibleRole looks for the Responsible Role cell in the control table.
func findResponsibleRole(ct *SummaryTable) (*responsibleRole, error) {
	nodes, err := ct.table.searchSubtree(".//w:tc[starts-with(normalize-space(.), 'Responsible Role')]")
	if err != nil {
		return nil, err
	}
	if len(nodes) != 1 {
		return nil, errors.New("could not find Responsible Role cell")
	}
	parentNode := nodes[0]
	childNodes, err := helper.SearchSubtree(parentNode, `.//w:t`)
	if err != nil || len(childNodes) < 1 {
		return nil, errors.New("should not happen, cannot find text nodes")
	}
	return &responsibleRole{parentNode: parentNode, textNodes: &childNodes}, nil
}

// responsibleRole is the container for the responsible role cell.
type responsibleRole struct {
	parentNode xml.Node
	textNodes  *[]xml.Node
}

// getContent returns the full string representation of the content of the cell itself.
func (r *responsibleRole) getContent() string {
	return r.parentNode.Content()
}

// setValue will set the value of the responsible role cell and do any needed formatting.
// In this case, it will just place the text after "Responsible Role: "
// If there are other nodes, we don't care about them, zero the content out.
func (r *responsibleRole) setValue(value string) {
	for idx, node := range *(r.textNodes) {
		if idx == 0 {
			node.SetContent(fmt.Sprintf("Responsible Role: %s", value))
		} else {
			node.SetContent("")
		}
	}
}

// isDefaultValue contains the logic to detect if the input is a default value. This is looking at the extracted
// value of responsible role and not the full string representation.
func (r *responsibleRole) isDefaultValue(value string) bool {
	return value == ""
}

// getValue extracts the unique value from the full string representation.
// It looks at all the text after "Responsible Role:".
func (r *responsibleRole) getValue() string {
	re := regexp.MustCompile("Responsible Role:?")
	result := ""
	// Get all the substrings
	subStrings := re.Split(r.parentNode.Content(), -1)
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
