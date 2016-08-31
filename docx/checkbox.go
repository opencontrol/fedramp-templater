package docx

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/docx/helper"
)

const checkBoxAttributeKey = "val"

// NewCheckBox constructs a new checkbox. Checks if the checkmark value can actually be found.
// If it cannot be found, will return nil.
func NewCheckBox(checkMark xml.Node, textNodes *[]xml.Node) *CheckBox {
	// Have to use Attr.
	// Using Attribute does not work when checking the value key.
	// Make sure the length is non zero.
	if len(checkMark.Attr(checkBoxAttributeKey)) == 0 {
		return nil
	}
	return &CheckBox{checkMark: checkMark, textNodes: textNodes}
}

// CheckBox represents a checkbox in a word document with any corresponding text.
type CheckBox struct {
	checkMark xml.Node
	textNodes *[]xml.Node
}

// IsChecked will return true if the box is checked, false otherwise.
func (c *CheckBox) IsChecked() bool {
	return c.checkMark.Attr(checkBoxAttributeKey) == "1"
}

// SetCheckMarkTo will set the checkbox state according to the input value.
func (c *CheckBox) SetCheckMarkTo(value bool) {
	checkBoxValue := "0"
	if value == true {
		checkBoxValue = "1"
	}
	panic(c.checkMark.AttributeList()[0])
	c.checkMark.AttributeList()[0].SetContent(checkBoxValue)
}

// GetTextValue will return the corresponding text for the checkbox.
func (c *CheckBox) GetTextValue() string {
	return helper.ConcatTextNodes(*(c.textNodes))
}