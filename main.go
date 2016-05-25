package main

import (
	"fmt"
	"os"

	"github.com/opencontrol/doc-template/docx"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage:\n\n\tfedramp-templater <input> <output>\n")
		os.Exit(1)
	}
	inputPath := os.Args[1]
	outputPath := os.Args[2]
	fmt.Printf("Creating template %s from %s...\n", outputPath, inputPath)

	doc := new(docx.Docx)
	doc.ReadFile("foo.doc")
}
