package templater_test

import (
	"io/ioutil"
	"path/filepath"

	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ControlTable", func() {
	Describe("Fill", func() {
		It("fills in the Responsible Role field", func() {
			path := filepath.Join("..", "fixtures", "simplified_table.xml")
			content, _ := ioutil.ReadFile(path)
			doc, _ := ParseWordXML(content)
			tables, _ := doc.Search("//w:tbl")
			table := tables[0]

			ct := ControlTable{Root: table}
			ct.Fill()

			Expect(table.Content()).To(ContainSubstring("Responsible Role: {{getResponsibleRole \"NIST-800-53\" \"AC-2 (1)\"}}"))
		})
	})
})
