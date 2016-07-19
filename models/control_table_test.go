package models_test

import (
	"bytes"
	"path/filepath"
	"text/template"
	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri/xml"

	. "github.com/opencontrol/fedramp-templater/models"
	"github.com/opencontrol/fedramp-templater/ssp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type tableData struct {
	Control string
}

func docFixture(control string) *xml.XmlDocument {
	path := filepath.Join("..", "fixtures", "simplified_table.xml")
	tpl, err := template.ParseFiles(path)
	Expect(err).ToNot(HaveOccurred())

	buf := new(bytes.Buffer)
	data := tableData{control}
	tpl.Execute(buf, data)

	doc, err := ssp.ParseWordXML(buf.Bytes())
	Expect(err).ToNot(HaveOccurred())

	return doc
}

var _ = Describe("ControlTable", func() {
	Describe("Fill", func() {
		It("fills in the Responsible Role for controls", func() {
			doc := docFixture("AC-2")
			tables, _ := doc.Search("//w:tbl")
			table := tables[0]

			ct := ControlTable{Root: table}
			ct.Fill()

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: {{getResponsibleRole "NIST-800-53" "AC-2"}}`))
		})

		It("fills in the Responsible Role for control enhancements", func() {
			doc := docFixture("AC-2 (1)")
			tables, _ := doc.Search("//w:tbl")
			table := tables[0]

			ct := ControlTable{Root: table}
			ct.Fill()

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: {{getResponsibleRole "NIST-800-53" "AC-2 (1)"}}`))
		})
	})
})
