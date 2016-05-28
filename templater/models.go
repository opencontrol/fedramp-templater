package templater

import (
	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri/xml"
)

type ControlTable struct {
  Root xml.Node
}

func (ct *ControlTable) ResponsibleRoleCell() (node xml.Node, err error) {
	nodes, err := ct.Root.Search("//w:tc//w:t[contains(., 'Responsible Role')]")
	if err != nil {
		return
	}
	node = nodes[0]
	return
}

// modifies the `table`
func (ct *ControlTable) Fill() (err error) {
	roleCell, err := ct.ResponsibleRoleCell()
	if err != nil {
		return
	}

	content := roleCell.Content()
	// TODO remove hard-coding
	content += " {{getResponsibleRole \"NIST-800-53\" \"AC-2 (1)\"}}"
	roleCell.SetContent(content)

	return
}
