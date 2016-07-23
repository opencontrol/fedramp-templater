package templater_test

import (
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"
	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func loadSSP(name string) *ssp.Document {
	sspPath := filepath.Join("..", "fixtures", name)
	doc, err := ssp.Load(sspPath)
	Expect(err).NotTo(HaveOccurred())

	return doc
}

func loadOpenControlFixture() opencontrols.Data {
	openControlDir := filepath.Join("..", "fixtures", "opencontrols")
	openControlDir, err := filepath.Abs(openControlDir)
	Expect(err).NotTo(HaveOccurred())
	openControlData, errors := opencontrols.LoadFrom(openControlDir)
	for _, err := range errors {
		Expect(err).NotTo(HaveOccurred())
	}

	return openControlData
}

var _ = Describe("Templater", func() {
	Describe("TemplatizeSSP", func() {
		It("fills in the Responsible Role fields", func() {
			doc := loadSSP("FedRAMP_ac-2-1_v2.1.docx")
			defer doc.Close()
			openControlData := loadOpenControlFixture()

			err := TemplatizeSSP(doc, openControlData)

			Expect(err).NotTo(HaveOccurred())
			content := doc.Content()
			Expect(content).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})
	})
})
