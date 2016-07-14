package templater

import (
	"fmt"

	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/doc-template/docx"
)

func GetWordDoc(path string) (doc *docx.Docx, err error) {
	doc = new(docx.Docx)
	err = doc.ReadFile(path)
	return
}

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

func getXMLDoc(wordDoc *docx.Docx) (xmlDoc *xml.XmlDocument, err error) {
	content := wordDoc.GetContent()
	// http://stackoverflow.com/a/28261008/358804
	bytes := []byte(content)
	return ParseWordXML(bytes)
}

// TODO merge these
func controlSummaryTables(doc *xml.XmlDocument) (tables []xml.Node, err error) {
	return doc.Search("//w:tbl[contains(., 'Control Summary')]")
}

func controlEnhancementSummaryTables(doc *xml.XmlDocument) (tables []xml.Node, err error) {
	return doc.Search("//w:tbl[contains(., 'Control Enhancement Summary')]")
}

// returns the tables for the controls and the control enhancements
func findSummaryTables(doc *xml.XmlDocument) (tables []xml.Node, err error) {
	controlTables, err := controlSummaryTables(doc)
	if err != nil {
		return
	}
	enhancementTables, err := controlEnhancementSummaryTables(doc)
	if err != nil {
		return
	}

	tables = append(controlTables, enhancementTables...)
	return
}

func printNode(node xml.Node) {
	fmt.Printf("%#v ", node)
	fmt.Printf("\"%s\"\n", node.Content())
}

func printNodes(nodes []xml.Node) {
	for _, node := range nodes {
		printNode(node)
	}
}

func templatizeXMLDoc(doc *xml.XmlDocument) (err error) {
	tables, err := findSummaryTables(doc)
	if err != nil {
		return
	}
	for _, table := range tables {
		ct := ControlTable{Root: table}
		err = ct.Fill()
		if err != nil {
			return err
		}
	}
	return
}

func TemplatizeWordDoc(wordDoc *docx.Docx) (err error) {
	xmlDoc, err := getXMLDoc(wordDoc)
	defer xmlDoc.Free()
	if err != nil {
		return
	}

	err = templatizeXMLDoc(xmlDoc)
	if err != nil {
		return
	}

	wordDoc.UpdateConent(xmlDoc.String())
	return
}
