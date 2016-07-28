package ssp_test

import (
	"path/filepath"

	. "github.com/opencontrol/fedramp-templater/ssp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func loadSSP(name string) *Document {
	path := filepath.Join("..", "fixtures", name)
	doc, err := Load(path)
	Expect(err).NotTo(HaveOccurred())
	return doc
}

var _ = Describe("SSP", func() {
	Describe("Load", func() {
		It("gets the content from the doc", func() {
			doc := loadSSP("FedRAMP_ac-2-1_v2.1.docx")
			defer doc.Close()

			Expect(doc.Content()).To(ContainSubstring("Control Enhancement"))
		})

		It("give an error when the doc isn't found", func() {
			_, err := Load("non-existent.docx")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("SummaryTables", func() {
		It("returns the tables", func() {
			doc := loadSSP("FedRAMP_ac-2_v2.1.docx")
			defer doc.Close()

			tables, err := doc.SummaryTables()

			Expect(err).NotTo(HaveOccurred())
			Expect(len(tables)).To(Equal(10))
		})
	})
})
