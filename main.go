package main

import (
	"fmt"
	"os"

	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/doc-template/docx"
)

func parseArgs() (inputPath, outputPath string) {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage:\n\n\tfedramp-templater <input> <output>\n\n")
		os.Exit(1)
	}
	inputPath = os.Args[1]
	outputPath = os.Args[2]
	return
}

func getWordDoc(path string) (doc *docx.Docx) {
	doc = new(docx.Docx)
	doc.ReadFile(path)
	return
}

func getXMLDoc(wordDoc *docx.Docx) (xmlDoc *xml.XmlDocument, err error) {
	content := wordDoc.GetContent()
	// http://stackoverflow.com/a/28261008/358804
	bytes := []byte(content)

	xmlDoc, err = gokogiri.ParseXml(bytes)
	if err != nil {
		return
	}
	// http://stackoverflow.com/a/27475227/358804
	xp := xmlDoc.DocXPathCtx()
	xp.RegisterNamespace("w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")

	return
}

func findControlEnhancementTables(doc *xml.XmlDocument) (nodes []xml.Node, err error) {
	return doc.Search("//w:tbl[contains(., 'Control Enhancement Summary')]")
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

func findResponsibleRoleCell(table xml.Node) (node xml.Node, err error) {
	nodes, err := table.Search("//w:tc//w:t[contains(., 'Responsible Role')]")
	if err != nil {
		return
	}
	node = nodes[0]
	return
}

func fillTable(table xml.Node) (err error) {
	roleCell, err := findResponsibleRoleCell(table)
	if err != nil {
		return
	}

	content := roleCell.Content()
	// TODO remove hard-coding
	content += " {{getResponsibleRole \"NIST-800-53\" \"AC-2 (1)\"}}"
	roleCell.SetContent(content)

	return
}

func templatizeXMLDoc(doc *xml.XmlDocument) (err error) {
	tables, err := findControlEnhancementTables(doc)
	if err != nil {
		return
	}
	for _, table := range tables {
		err = fillTable(table)
		if err != nil {
			return err
		}
	}

	return
}

func templatizeWordDoc(wordDoc *docx.Docx) (err error) {
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

func main() {
	inputPath, outputPath := parseArgs()
	wordDoc := getWordDoc(inputPath)

	err := templatizeWordDoc(wordDoc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// TODO this should use the current content, or not be a method
	wordDoc.WriteToFile(outputPath, wordDoc.GetContent())
}
