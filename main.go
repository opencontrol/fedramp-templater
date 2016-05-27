package main

import (
	"fmt"
	"os"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
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

func getXmlDoc(path string) (xmlDoc *xml.XmlDocument, err error) {
	wordDoc := new(docx.Docx)
	wordDoc.ReadFile(path)

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
	query := xpath.Compile("//w:tbl[contains(., 'Control Enhancement Summary')]")
	return doc.Search(query)
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
	query := xpath.Compile("//w:tc[contains(., 'Responsible Role')]")
	nodes, err := table.Search(query)
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
	printNode(roleCell)

	return
}

func handleXmlDoc(doc *xml.XmlDocument) (err error) {
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

func main() {
	// inputPath, outputPath := parseArgs()
	inputPath, _ := parseArgs()

	doc, err := getXmlDoc(inputPath)
	defer doc.Free()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = handleXmlDoc(doc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// doc.WriteToFile(outputPath, content)
}
