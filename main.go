package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/templater"
)

func parseArgs() (inputPath, outputPath string) {
	if len(os.Args) != 3 {
		log.Fatal("Usage:\n\n\tfedramp-templater <input> <output>\n\n")
	}
	inputPath = os.Args[1]
	outputPath = os.Args[2]
	return
}

func main() {
	inputPath, outputPath := parseArgs()
	wordDoc, err := templater.GetWordDoc(inputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = templater.TemplatizeWordDoc(wordDoc)
	if err != nil {
		log.Fatalln(err)
	}

	outputDir := filepath.Dir(outputPath)
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	// TODO this should use the current content, or not be a method
	wordDoc.WriteToFile(outputPath, wordDoc.GetContent())
}
