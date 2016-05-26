package main

import (
	"fmt"
	"os"

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
	inputPath, outputPath := parseArgs()

	doc := new(docx.Docx)
	doc.ReadFile(inputPath)
	doc.WriteToFile(outputPath, doc.GetContent())
}
