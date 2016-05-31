package templater_test

import (
	"io/ioutil"
	"path/filepath"

	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Templater", func() {
	Describe("FillTable", func() {
		It("fills in the Responsible Role field", func() {
			path := filepath.Join("..", "fixtures", "simplified_table.xml")
			content, _ := ioutil.ReadFile(path)
			doc, _ := ParseWordXML(content)
			tables, _ := doc.Search("//w:tbl")
			table := tables[0]

			FillTable(table)

			Expect(table.Content()).To(ContainSubstring("getResponsibleRole"))
		})
	})

	Describe("GetWordDoc", func() {
		It("gets the content from the doc", func() {
			path := filepath.Join("..", "fixtures", "FedRAMP_ac-2-1_v2.1.docx")
			doc := GetWordDoc(path)

			Expect(doc.GetContent()).To(ContainSubstring("Control Enhancement"))
		})

		It("returns an empty string when the doc isn't found", func() {
			doc := GetWordDoc("non-existent.docx")
			Expect(doc.GetContent()).To(Equal(""))
		})
	})
})
