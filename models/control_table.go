package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri/xml"
)

// ControlTable represents the node in the Word docx XML tree that corresponds to a security control.
type ControlTable struct {
	Root xml.Node
}

func (ct *ControlTable) searchSubtree(xpath string) (nodes []xml.Node, err error) {
	// http://stackoverflow.com/a/25387687/358804
	if !strings.HasPrefix(xpath, ".") {
		err = errors.New("XPath must have leading period (`.`) to only search the subtree")
		return
	}

	return ct.Root.Search(xpath)
}

func (ct *ControlTable) responsibleRoleCell() (node xml.Node, err error) {
	nodes, err := ct.searchSubtree(".//w:tc//w:t[contains(., 'Responsible Role')]")
	if err != nil {
		return
	}
	if len(nodes) != 1 {
		err = errors.New("could not find Responsible Role cell")
		return
	}
	node = nodes[0]
	return
}

func (ct *ControlTable) tableHeader() (content string, err error) {
	nodes, err := ct.searchSubtree(".//w:tr")
	if err != nil {
		return
	}
	if len(nodes) == 0 {
		err = errors.New("could not find control name")
		return
	}
	// we only care about the first match
	content = nodes[0].Content()

	return
}

func (ct *ControlTable) controlName() (name string, err error) {
	content, err := ct.tableHeader()
	if err != nil {
		return
	}

	// matches controls and control enhancements, e.g. `AC-2`, `AC-2 (1)`, etc.
	regex := regexp.MustCompile(`[A-Z]{2}-\d+( +\(\d+\))?`)
	name = regex.FindString(content)
	if name == "" {
		err = errors.New("control name not found")
	}
	return
}

// Fill inserts the tags into the table. Note this modifies the `table`.
func (ct *ControlTable) Fill() (err error) {
	roleCell, err := ct.responsibleRoleCell()
	if err != nil {
		return
	}

	existingContent := roleCell.Content()
	standard := "NIST-800-53"
	control, err := ct.controlName()
	if err != nil {
		return
	}

	content := fmt.Sprintf("%s {{getResponsibleRole %q %q}}", existingContent, standard, control)
	roleCell.SetContent(content)

	return
}
