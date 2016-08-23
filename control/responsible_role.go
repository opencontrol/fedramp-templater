package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"errors"
	"regexp"
	"strings"
)

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

type ResponsibleRole struct {
	node xml.Node
}

func (r *ResponsibleRole) GetContent() string {
	return r.node.Content()
}

func (r *ResponsibleRole) SetContent(content string) {
	r.node.SetContent(content)
}


func (r *ResponsibleRole) IsDefaultValue(value string) bool {
	return value == ""
}

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