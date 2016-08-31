package docx

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/docx/helper"
)

const checkBoxAttributeKey = "val"

func NewCheckBox(checkMark xml.Node, textNodes *[]xml.Node) *CheckBox {
	// Have to use Attr.
	// Using Attribute does not work when checking the value key.
	// Make sure the length is non zero.
	if len(checkMark.Attr(checkBoxAttributeKey)) == 0 {
		return nil
	}
	return &CheckBox{checkMark: checkMark, textNodes: textNodes}
}

type CheckBox struct {
	checkMark xml.Node
	textNodes *[]xml.Node
}

func (c *CheckBox) IsChecked() bool {
	return c.checkMark.Attr(checkBoxAttributeKey) == "1"
}

func (c *CheckBox) SetCheckMarkTo(value bool) {
	checkBoxValue := "0"
	if value == true {
		checkBoxValue = "1"
	}
	panic(c.checkMark.AttributeList()[0])
	c.checkMark.AttributeList()[0].SetContent(checkBoxValue)
}

func (c *CheckBox) GetTextValue() string {
	return helper.ConcatTextNodes(*(c.textNodes))
}