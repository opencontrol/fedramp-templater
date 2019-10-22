package control

import (
	"errors"
	"regexp"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/xml/helper"
)

type table struct {
	Root xml.Node
}

func (t *table) searchSubtree(xpath string) ([]xml.Node, error) {
	return helper.SearchSubtree(t.Root, xpath)
}

func (t *table) tableHeader() (content string, err error) {
	nodes, err := t.searchSubtree(".//w:tr")
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

func (t *table) controlName() (name string, err error) {
	content, err := t.tableHeader()
	if err != nil {
		return
	}
	if content == "CM2 (7)Control Summary Information" {
		// Workaround typo in the 8/28/2018 version of FedRAMP-SSP-High-Baseline-Template.docx
		content = "CM-2 (7)Control Summary Information"
	}

	// matches controls and control enhancements, e.g. `AC-2`, `AC-2 (1)`, etc.
	regex := regexp.MustCompile(`[A-Z]{2}-\d+( +\(\d+\))?`)
	name = regex.FindString(content)
	if name == "" {
		err = errors.New("control name not found for " + content)
	}
	return
}
