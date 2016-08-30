package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"strings"
)

func newCheckBox(checkMark xml.Node, textNodes *[]xml.Node) *checkBox {
	return &checkBox{checkMark: checkMark, textNodes: textNodes}
}

type checkBox struct {
	checkMark xml.Node
	textNodes *[]xml.Node
}

func (c *checkBox) isChecked() bool {
	return false
}

func (c *checkBox) setCheckMarkTo(value bool) {

}

func (c *checkBox) getTextValue() string {
	return concatTextNodes(*(c.textNodes))
}


func concatTextNodes(textNodes []xml.Node) string {
	result := ""
	for _, textNode := range textNodes {
		result += textNode.Content()
	}
	return strings.TrimSpace(result)
}