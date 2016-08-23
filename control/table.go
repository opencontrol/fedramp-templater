package control

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/opencontrols"
)

// Table represents the node in the Word docx XML tree that corresponds to a security control.
type Table struct {
	Root xml.Node
}

func (ct *Table) searchSubtree(xpath string) (nodes []xml.Node, err error) {
	// http://stackoverflow.com/a/25387687/358804
	if !strings.HasPrefix(xpath, ".") {
		err = errors.New("XPath must have leading period (`.`) to only search the subtree")
		return
	}

	return ct.Root.Search(xpath)
}

func (ct *Table) tableHeader() (content string, err error) {
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

func (ct *Table) controlName() (name string, err error) {
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

// Fill inserts the OpenControl justifications into the table. Note this modifies the `table`.
func (ct *Table) Fill(openControlData opencontrols.Data) (err error) {
	roleCell, err := FindResponsibleRole(ct)
	if err != nil {
		return
	}

	existingContent := roleCell.GetContent()
	control, err := ct.controlName()
	if err != nil {
		return
	}

	roles := openControlData.GetResponsibleRoles(control)
	content := fmt.Sprintf("%s %s", existingContent, roles)
	roleCell.SetContent(content)

	return
}

func (ct *Table) Diff(openControlData opencontrols.Data) string {
	control, err := ct.controlName()
	if err != nil {
		return ""
	}

	roleCell, err := FindResponsibleRole(ct)
	if err != nil {
		return ""
	}
	roles := openControlData.GetResponsibleRoles(control)
	value := roleCell.GetValue()
	if roleCell.IsDefaultValue(value) || roles == value {
		return ""
	}
	return fmt.Sprintf("Control: %s. Responsible Role in doc :\"%s\". Responsible Role in YAML \"%s\"\n", control, value, roles)
}