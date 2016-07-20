package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/ssp"
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

	doc, err := ssp.Load(inputPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer doc.Close()

	err = templater.TemplatizeSSP(doc)
	if err != nil {
		log.Fatalln(err)
	}

	outputDir := filepath.Dir(outputPath)
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalln(err)
	}

	err = doc.CopyTo(outputPath)
	if err != nil {
		log.Fatalln(err)
	}
}
