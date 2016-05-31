package templater

import (
	"fmt"

	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri/xml"
)

type ControlTable struct {
  Root xml.Node
}

func (ct *ControlTable) responsibleRoleCell() (node xml.Node, err error) {
	nodes, err := ct.Root.Search("//w:tc//w:t[contains(., 'Responsible Role')]")
	if err != nil {
		return
	}
	node = nodes[0]
	return
}

func (ct *ControlTable) controlName() (name string) {
	// TODO remove hard-coding
	return "AC-2 (1)"
}

// modifies the `table`
func (ct *ControlTable) Fill() (err error) {
	roleCell, err := ct.responsibleRoleCell()
	if err != nil {
		return
	}

	existingContent := roleCell.Content()
	standard := "NIST-800-53"
	content := fmt.Sprintf("%s {{getResponsibleRole %q %q}}", existingContent, standard, ct.controlName())
	roleCell.SetContent(content)

	return
}
