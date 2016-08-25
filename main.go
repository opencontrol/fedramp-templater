package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"
	"github.com/opencontrol/fedramp-templater/templater"
	"flag"
)

type options struct {
	openControlsDir string
	inputPath string
	outputPath string
	diff bool
	fill bool
}

func parseArgs() (opts options) {
	flag.Usage = func() {
		log.Fatal("Usage:\n\n" +
			"\tfedramp-templater -fill <openControlsDir> <inputDoc> <outputDoc>\n\n" +
			"\tor\n\n" +
			"\tfedramp-templater -diff <openControlsDir> <inputDoc>")
	}
	diffPtr := flag.Bool("diff", false, "flag to indicate using diff.")
	fillPtr := flag.Bool("fill", false, "flag to indicate using fill.")
	flag.Parse()
	if (*diffPtr && flag.NArg() == 2 && !*fillPtr) {
		// diff command only has two args
		opts.diff = *diffPtr
		opts.openControlsDir = flag.Args()[0]
		opts.inputPath = flag.Args()[1]
	} else if (*fillPtr && flag.NArg() == 3 && !*diffPtr) {
		// fill command only has two args
		opts.fill = *fillPtr
		opts.openControlsDir = flag.Args()[0]
		opts.inputPath = flag.Args()[1]
		opts.outputPath = flag.Args()[2]
	} else {
		flag.Usage()
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
