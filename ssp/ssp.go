package ssp

import (
	"github.com/opencontrol/doc-template/docx"
	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
)

// ParseWordXML converts the XML text to a structure.
func ParseWordXML(content []byte) (xmlDoc *xml.XmlDocument, err error) {
	xmlDoc, err = gokogiri.ParseXml(content)
	if err != nil {
		return
	}
	// http://stackoverflow.com/a/27475227/358804
	xp := xmlDoc.DocXPathCtx()
	xp.RegisterNamespace("w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")
	return
}

func generateXml(wordDoc *docx.Docx) (xmlDoc *xml.XmlDocument, err error) {
	content := wordDoc.GetContent()
	// http://stackoverflow.com/a/28261008/358804
	bytes := []byte(content)
	return ParseWordXML(bytes)
}

type Ssp struct {
	wordDoc *docx.Docx
	xmlDoc  *xml.XmlDocument
}

func getWordDoc(path string) (doc *docx.Docx, err error) {
	doc = new(docx.Docx)
	err = doc.ReadFile(path)
	return
}

func Load(path string) (ssp *Ssp, err error) {
	doc, err := getWordDoc(path)
	if err != nil {
		return
	}
	xmlDoc, err := generateXml(doc)
	if err != nil {
		return
	}

	ssp = &Ssp{
		wordDoc: doc,
		xmlDoc:  xmlDoc,
	}
	return
}

// SummaryTables returns the tables for the controls and the control enhancements.
func (s *Ssp) SummaryTables() (tables []xml.Node, err error) {
	// find the tables matching the provided headers, ignoring whitespace
	return s.xmlDoc.Search("//w:tbl[contains(normalize-space(.), 'Control Summary') or contains(normalize-space(.), 'Control Enhancement Summary')]")
}

func (s *Ssp) Content() string {
	return s.wordDoc.GetContent()
}

func (s *Ssp) UpdateContent() {
	content := s.xmlDoc.String()
	// TODO fix spelling upstream
	s.wordDoc.UpdateConent(content)
}

func (s *Ssp) CopyTo(path string) {
	// TODO fix upstream: WriteToFile should use the doc's content, or not be a method
	s.wordDoc.WriteToFile(path, s.Content())
}

func (s *Ssp) Close() error {
	s.xmlDoc.Free()
	return s.wordDoc.Close()
}
