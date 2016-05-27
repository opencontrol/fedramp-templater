package main

import (
	"fmt"
	"os"

	"github.com/opencontrol/fedramp-templater/templater"
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
	wordDoc := templater.GetWordDoc(inputPath)

	err := templater.TemplatizeWordDoc(wordDoc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// TODO this should use the current content, or not be a method
	wordDoc.WriteToFile(outputPath, wordDoc.GetContent())
}
