package fixtures

import (
	"os"
	"path/filepath"

	dbg "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"
)

// FixturePath - path of the fixture
func FixturePath(name string) string {
	path := filepath.Join("..", "fixtures", name)
	path, err := filepath.Abs(path)
	Expect(err).NotTo(HaveOccurred())
	// ensure the file/directory exists
	_, err = os.Stat(path)
	Expect(err).NotTo(HaveOccurred())

	return path
}

// OpenControlFixturePath - Path of the OpenControl fixture
func OpenControlFixturePath() string {
	return FixturePath("opencontrols")
}

// LoadSSP - load an SSP document
func LoadSSP(name string) *ssp.Document {
	sspPath := FixturePath(name)
	doc, err := ssp.Load(sspPath)
	Expect(err).NotTo(HaveOccurred())

	return doc
}

// LoadOpenControlFixture - Load an OpenControl fixture
func LoadOpenControlFixture() opencontrols.Data {
	openControlDir := OpenControlFixturePath()
	openControlData, errors := opencontrols.LoadFrom(openControlDir)
	for _, err := range errors {
		Expect(err).NotTo(HaveOccurred())
	}

	return openControlData
}
