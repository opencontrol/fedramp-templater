package templater_test

import (
	"path/filepath"

	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Templater", func() {
	Describe("GetWordDoc", func() {
		It("gets the content from the doc", func() {
			path := filepath.Join("..", "fixtures", "FedRAMP_ac-2-1_v2.1.docx")
			doc, err := GetWordDoc(path)

			Expect(err).NotTo(HaveOccurred())
			Expect(doc.GetContent()).To(ContainSubstring("Control Enhancement"))
		})

		It("give an error when the doc isn't found", func() {
			doc, err := GetWordDoc("non-existent.docx")
			Expect(err).To(HaveOccurred())
			Expect(doc.GetContent()).To(Equal(""))
		})
	})
})
