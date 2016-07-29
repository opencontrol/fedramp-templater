package opencontrols_test

import (
	"path/filepath"

	. "github.com/opencontrol/fedramp-templater/opencontrols"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TODO share with templater_test
func loadOpenControlFixture() Data {
	openControlDir := filepath.Join("..", "fixtures", "opencontrols")
	openControlDir, err := filepath.Abs(openControlDir)
	Expect(err).NotTo(HaveOccurred())
	openControlData, errors := LoadFrom(openControlDir)
	for _, err := range errors {
		Expect(err).NotTo(HaveOccurred())
	}

	return openControlData
}

var _ = Describe("Data", func() {
	Describe("GetNarrative", func() {
		It("returns the relevant singular narrative", func() {
			data := loadOpenControlFixture()
			result := data.GetNarrative("AC-2 (1)", "")
			Expect(result).To(Equal("Amazon Elastic Compute Cloud\nJustification in narrative form for AC-2 (1)\n"))
		})

		It("returns the relevant narrative section", func() {
			data := loadOpenControlFixture()
			result := data.GetNarrative("AC-2", "a")
			Expect(result).To(Equal("Amazon Elastic Compute Cloud\nJustification in narrative form A for AC-2\n"))
		})
	})
})
