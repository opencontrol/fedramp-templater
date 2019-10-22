package helper

import (
	"strings"

	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/doc-template/docx"
)

// ParseXML converts the XML text to a structure.
func ParseXML(content []byte) (xmlDoc *xml.XmlDocument, err error) {
	xmlDoc, err = gokogiri.ParseXml(content)
	if err != nil {
		return
	}
	// http://stackoverflow.com/a/27475227/358804
	xp := xmlDoc.DocXPathCtx()
	xp.RegisterNamespace("w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")
	xp.RegisterNamespace("w14", "http://schemas.microsoft.com/office/word/2010/wordml")
	return
}

// GenerateXML gives the underlying XML document for the provided Word document.
func GenerateXML(wordDoc *docx.Docx) (xmlDoc *xml.XmlDocument, err error) {
	content := wordDoc.GetContent()
	// http://stackoverflow.com/a/28261008/358804
	bytes := []byte(content)
	return ParseXML(bytes)
}

// FillParagraph inserts the given content into the provided docx XML paragraph node. Note that newlines aren't respected - you'll need to create a new paragraph node for each.
func FillParagraph(paragraph xml.Node, content string) (err error) {
	// this seems to be the easiest way to create child notes
	err = paragraph.SetChildren(`<w:r><w:t></w:t></w:r>`)
	if err != nil {
		return
	}
	textCell := paragraph.FirstChild().FirstChild()

	textCell.SetContent(content)
	return
}

// AddMultiLineContent adds the given content into the provided docx XML node as multiple "paragraphs" of text, split by the newlines.
func AddMultiLineContent(parent xml.Node, content string) error {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		// this seems to be the easiest way to create child notes
		err := parent.AddChild(`<w:p></w:p>`)
		if err != nil {
			return err
		}
		paragraph := parent.LastChild()
		FillParagraph(paragraph, line)
	}

	return nil
}

// FillCell inserts the given content into the provided docx XML table cell node.
func FillCell(cell xml.Node, content string) error {
	// clear out the existing content/structure
	err := cell.SetChildren("")
	if err != nil {
		return err
	}

	return AddMultiLineContent(cell, content)
}

// ConcatTextNodes will concatenate the text from an array of text nodes and trim any whitespace from the final result.
func ConcatTextNodes(textNodes []xml.Node) string {
	result := ""
	for _, textNode := range textNodes {
		result += textNode.Content()
	}
	return strings.TrimSpace(result)
}
