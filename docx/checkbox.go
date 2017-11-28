package docx

import (
	"fmt"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/docx/helper"
)

const (
	checkBoxAttributeKey    = "val"
	checkBoxCheckedValue    = "1"
	checkBoxNotCheckedValue = "0"
)

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
	return c.checkMark.Attr(checkBoxAttributeKey) == checkBoxCheckedValue
}

// SetCheckMarkTo will set the checkbox state according to the input value.
func (c *CheckBox) SetCheckMarkTo(value bool) {
	checkBoxValue := checkBoxNotCheckedValue
	if value == true {
		checkBoxValue = checkBoxCheckedValue
	}
	c.checkMark.AttributeList()[0].SetContent(checkBoxValue)
}

// GetTextValue will return the corresponding text for the checkbox.
func (c *CheckBox) GetTextValue() string {
	return helper.ConcatTextNodes(*(c.textNodes))
}

// FindCheckBoxTag will look for a checkbox and the value tag.
// We look for the w:default tag embedded in the w:checkBox tag because that is what we need to modify the checkbox.
func FindCheckBoxTag(paragraph xml.Node) (xml.Node, error) {
	checkBoxes, err := paragraph.Search(".//w:checkBox//w:default")
	if err != nil {
		return nil, err
	} else if len(checkBoxes) != 1 {
		return nil, fmt.Errorf("unable to find the check box value")
	}
	return checkBoxes[0], nil
}
