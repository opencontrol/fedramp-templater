package helper

import (
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
	return
}

// GenerateXML gives the underlying XML document for the provided Word document.
func GenerateXML(wordDoc *docx.Docx) (xmlDoc *xml.XmlDocument, err error) {
	content := wordDoc.GetContent()
	// http://stackoverflow.com/a/28261008/358804
	bytes := []byte(content)
	return ParseXML(bytes)
}

// FillParagraph inserts the given content into the provided docx XML paragraph node.
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
