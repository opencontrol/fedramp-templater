package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"strings"
)

const checkBoxAttributeKey = "val"

func newCheckBox(checkMark xml.Node, textNodes *[]xml.Node) *checkBox {
	return &checkBox{checkMark: checkMark, textNodes: textNodes}
}

type checkBox struct {
	checkMark xml.Node
	textNodes *[]xml.Node
}

func (c *checkBox) isChecked() bool {
	if c.checkMark.Attr(checkBoxAttributeKey) == "1" {
		return true
	}
	return false
}

func (c *checkBox) setCheckMarkTo(value bool) {
	checkBoxValue := "0"
	if value == true {
		checkBoxValue = "1"
	}
	c.checkMark.SetAttr(checkBoxAttributeKey, checkBoxValue)
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