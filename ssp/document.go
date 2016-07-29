package ssp

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/doc-template/docx"
	"github.com/opencontrol/fedramp-templater/docx/helper"
)

// Document represents a system security plan file and its contents.
type Document struct {
	wordDoc *docx.Docx
	xmlDoc  *xml.XmlDocument
}

func getWordDoc(path string) (doc *docx.Docx, err error) {
	doc = new(docx.Docx)
	err = doc.ReadFile(path)
	return
}

// Load creates a new Document from the provided file path.
func Load(path string) (ssp *Document, err error) {
	wordDoc, err := getWordDoc(path)
	if err != nil {
		return
	}
	xmlDoc, err := helper.GenerateXML(wordDoc)
	if err != nil {
		return
	}

	ssp = &Document{wordDoc, xmlDoc}
	return
}

// SummaryTables returns the tables for the controls and the control enhancements.
func (s *Document) SummaryTables() (tables []xml.Node, err error) {
	// find the tables matching the provided headers, ignoring whitespace
	return s.xmlDoc.Search("//w:tbl[contains(normalize-space(.), 'Control Summary') or contains(normalize-space(.), 'Control Enhancement Summary')]")
}

// Content retrieves the text from within the Word document.
func (s *Document) Content() string {
	return s.wordDoc.GetContent()
}

// UpdateContent modifies the state of the underlying Word document. Note this is purely for bookkeeping in memory, and does not actually make any changes to the file.
func (s *Document) UpdateContent() {
	content := s.xmlDoc.String()
	s.wordDoc.UpdateContent(content)
}

// CopyTo copies the contents of this Word document to a new file at the provided path.
func (s *Document) CopyTo(path string) error {
	return s.wordDoc.WriteToFile(path, s.Content())
}

// Close releases the underlying resources.
func (s *Document) Close() error {
	s.xmlDoc.Free()
	return s.wordDoc.Close()
}
