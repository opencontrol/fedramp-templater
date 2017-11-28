package ssp

import (
	"errors"
	"fmt"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/doc-template/docx"
	"github.com/opencontrol/fedramp-templater/docx/helper"
)

// SummaryTablesXPath is the pattern used to find summary tables within an SSP's XML.
const SummaryTablesXPath = "//w:tbl[contains(normalize-space(.), 'Control Summary') or contains(normalize-space(.), 'Control Enhancement Summary')]"

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
	if ssp != nil || err != nil {
		return
	}
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

// SummaryTables returns the summary tables for the controls and the control enhancements.
func (s *Document) SummaryTables() ([]xml.Node, error) {
	// find the tables matching the provided headers, ignoring whitespace
	return s.xmlDoc.Search(SummaryTablesXPath)
}

// to retrieve all narrative tables, pass in an empty string
func (s *Document) findNarrativeTables(control string) ([]xml.Node, error) {
	// find the tables matching the provided headers, ignoring whitespace
	xpath := fmt.Sprintf("//w:tbl[contains(normalize-space(.), '%s What is the solution and how is it implemented?')]", control)
	return s.xmlDoc.Search(xpath)
}

// NarrativeTables returns the narrative tables for all controls and the control enhancements.
func (s *Document) NarrativeTables() ([]xml.Node, error) {
	return s.findNarrativeTables("")
}

// NarrativeTable returns the narrative table for the specified control or control enhancement.
func (s *Document) NarrativeTable(control string) (table xml.Node, err error) {
	tables, err := s.findNarrativeTables(control)
	if err != nil {
		return
	}
	if len(tables) == 0 {
		err = errors.New("no narrative tables found")
		return
	} else if len(tables) > 1 {
		err = errors.New("too many narrative tables were matched")
		return
	}
	table = tables[0]
	return
}

// to retrieve all parameter tables, pass in an empty string
func (s *Document) findParameterTables(control string) ([]xml.Node, error) {
	// find the tables matching the provided headers, ignoring whitespace
	xpath := fmt.Sprintf("//w:tbl[contains(normalize-space(), 'Control Summary') or contains(normalize-space(.), 'Control Enhancement Summary')]")
	return s.xmlDoc.Search(xpath)
}

// ParameterTables returns the parameter tables for all controls and the control enhancements.
func (s *Document) ParameterTables() ([]xml.Node, error) {
	return s.findParameterTables("")
}

// ParameterTable returns the parameter table for the specified control or control enhancement.
func (s *Document) ParameterTable(control string) (table xml.Node, err error) {
	tables, err := s.findParameterTables(control)
	if err != nil {
		return
	}
	if len(tables) == 0 {
		err = errors.New("no parameter tables found")
		return
	} else if len(tables) > 1 {
		err = errors.New("too many parameter tables were matched")
		return
	}
	table = tables[0]
	return
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
