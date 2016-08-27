package fixtures

import (
	"os"
	"path/filepath"

	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"
)

func FixturePath(name string) string {
	path := filepath.Join("..", "fixtures", name)
	path, err := filepath.Abs(path)
	Expect(err).NotTo(HaveOccurred())
	// ensure the file/directory exists
	_, err = os.Stat(path)
	Expect(err).NotTo(HaveOccurred())

	return path
}

func OpenControlFixturePath() string {
	return FixturePath("opencontrols")
}

func LoadSSP(name string) *ssp.Document {
	sspPath := FixturePath(name)
	doc, err := ssp.Load(sspPath)
	Expect(err).NotTo(HaveOccurred())

	return doc
}

func LoadOpenControlFixture() opencontrols.Data {
	openControlDir := OpenControlFixturePath()
	openControlData, errors := opencontrols.LoadFrom(openControlDir)
	for _, err := range errors {
		Expect(err).NotTo(HaveOccurred())
	}

	return openControlData
}
