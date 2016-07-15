package templater_test

import (
	"bytes"
	"path/filepath"
	"text/template"

	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type tableData struct {
	Control string
}

var _ = Describe("ControlTable", func() {
	Describe("Fill", func() {
		It("fills in the Responsible Role field", func() {
			path := filepath.Join("..", "fixtures", "simplified_table.xml")
			tpl, err := template.ParseFiles(path)
			Expect(err).ToNot(HaveOccurred())

			buf := new(bytes.Buffer)
			data := tableData{"AC-2 (1)"}
			tpl.Execute(buf, data)

			doc, _ := ParseWordXML(buf.Bytes())
			tables, _ := doc.Search("//w:tbl")
			table := tables[0]

			ct := ControlTable{Root: table}
			ct.Fill()

			Expect(table.Content()).To(ContainSubstring("Responsible Role: {{getResponsibleRole \"NIST-800-53\" \"AC-2 (1)\"}}"))
		})
	})
})
