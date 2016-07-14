package templater

import (
	"errors"
	"fmt"
	"regexp"

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

func (ct *ControlTable) tableHeader() (content string, err error) {
	nodes, err := ct.Root.Search("//w:tr")
	if err != nil {
		return
	}
	if len(nodes) == 0 {
		err = errors.New("Could not find control name.")
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

	regex := regexp.MustCompile(`[A-Z]{2}-\d+( \(.\))?`)
	name = regex.FindString(content)
	if name == "" {
		err = errors.New("Control name not found.")
	}
	return
}

// modifies the `table`
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
