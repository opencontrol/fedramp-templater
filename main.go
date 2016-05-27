package main

import (
	"fmt"
	"os"

	"github.com/moovweb/gokogiri"
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

func main() {
	// inputPath, outputPath := parseArgs()
	inputPath, _ := parseArgs()

	wordDoc := new(docx.Docx)
	wordDoc.ReadFile(inputPath)

	content := wordDoc.GetContent()
	// http://stackoverflow.com/a/28261008/358804
	bytes := []byte(content)

	xmlDoc, err := gokogiri.ParseXml(bytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer xmlDoc.Free()

	// http://stackoverflow.com/a/27475227/358804
	xp := xmlDoc.DocXPathCtx()
	xp.RegisterNamespace("w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")

	query := xpath.Compile("//w:tbl")
	nodes, err := xmlDoc.Search(query)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(len(nodes))

	// doc.WriteToFile(outputPath, content)
}
