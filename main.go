package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"
	"github.com/opencontrol/fedramp-templater/templater"
)

func parseArgs() (openControlsDir, inputPath, outputPath string) {
	if len(os.Args) != 4 {
		log.Fatal("Usage:\n\n\tfedramp-templater <openControlsDir> <inputDoc> <outputDoc>\n\n")
	}
	openControlsDir = os.Args[1]
	inputPath = os.Args[2]
	outputPath = os.Args[3]
	return
}

func loadOpenControls(path string) opencontrols.Data {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
	}

	openControlData, errors := opencontrols.LoadFrom(path)
	if len(errors) > 0 {
		log.Fatal(errors)
	}
	return openControlData
}

func main() {
	openControlsDir, inputPath, outputPath := parseArgs()

	openControlData := loadOpenControls(openControlsDir)

	doc, err := ssp.Load(inputPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer doc.Close()

	err = templater.TemplatizeSSP(doc, openControlData)
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
