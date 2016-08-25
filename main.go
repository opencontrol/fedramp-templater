package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"
	"github.com/opencontrol/fedramp-templater/templater"
)

type options struct {
	openControlsDir string
	inputPath string
	outputPath string
	diff bool
	fill bool
}

func printUsage() {
	log.Fatal("Usage:\n\n" +
		"\tfedramp-templater fill <openControlsDir> <inputDoc> <outputDoc>\n\n" +
		"\tor\n\n" +
		"\tfedramp-templater diff <openControlsDir> <inputDoc>")
}

func parseArgs() (opts options) {
	if len(os.Args) < 4 || len(os.Args) > 5 {
		printUsage()
	}
	switch (os.Args[1]) {
	case "diff":
		opts.diff = true
	case "fill":
		opts.fill = true
	default:
		log.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
	}
	if (opts.diff && len(os.Args) == 4) {
		// diff command only has four args
		opts.openControlsDir = os.Args[2]
		opts.inputPath = os.Args[3]
	} else if (opts.fill && len(os.Args) == 5) {
		// fill command only has five args
		opts.openControlsDir = os.Args[2]
		opts.inputPath = os.Args[3]
		opts.outputPath = os.Args[4]
	} else {
		printUsage()
	}
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

func diffCmd(openControlData opencontrols.Data, doc *ssp.Document) {
	reporters, err := templater.DiffSSP(doc, openControlData)
	if err != nil {
		log.Fatalln(err)
	}
	if len(reporters) == 0 {
		log.Println("No diff detected")
		return
	}
	for _, reporter := range reporters{
		reporter.WriteTextTo(os.Stdout)
	}
	log.Fatalf("%d diffs detected\n", len(reporters))
}

func fillCmd(openControlData opencontrols.Data, doc *ssp.Document, opts options) {
	err := templater.TemplatizeSSP(doc, openControlData)
	if err != nil {
		log.Fatalln(err)
	}

	outputDir := filepath.Dir(opts.outputPath)
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalln(err)
	}

	err = doc.CopyTo(opts.outputPath)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	opts := parseArgs()

	openControlData := loadOpenControls(opts.openControlsDir)

	doc, err := ssp.Load(opts.inputPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer doc.Close()

	// right now we don't want to do a fill and diff together.
	if opts.diff {
		diffCmd(openControlData, doc)

	} else if opts.fill {
		fillCmd(openControlData, doc, opts)
	}
}
