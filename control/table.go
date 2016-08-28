package control

import (
	"errors"
	"regexp"
	"strings"

	"github.com/jbowtie/gokogiri/xml"
)

type table struct {
	Root xml.Node
}

func (t *table) searchSubtree(xpath string) (nodes []xml.Node, err error) {
	// http://stackoverflow.com/a/25387687/358804
	if !strings.HasPrefix(xpath, ".") {
		err = errors.New("XPath must have leading period (`.`) to only search the subtree")
		return
	}

	return t.Root.Search(xpath)
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

	// matches controls and control enhancements, e.g. `AC-2`, `AC-2 (1)`, etc.
	regex := regexp.MustCompile(`[A-Z]{2}-\d+( +\(\d+\))?`)
	name = regex.FindString(content)
	if name == "" {
		err = errors.New("control name not found")
	}
	return
}
