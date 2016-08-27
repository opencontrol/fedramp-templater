package control_test

import (
	"bytes"
	"text/template"

	"github.com/jbowtie/gokogiri/xml"
	. "github.com/opencontrol/fedramp-templater/control"
	"github.com/opencontrol/fedramp-templater/docx/helper"
	"github.com/opencontrol/fedramp-templater/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type tableData struct {
	Control string
}

func docFixture(control string) *xml.XmlDocument {
	path := fixtures.FixturePath("simplified_table.xml")
	tpl, err := template.ParseFiles(path)
	Expect(err).ToNot(HaveOccurred())

	buf := new(bytes.Buffer)
	data := tableData{control}
	tpl.Execute(buf, data)

	doc, err := helper.ParseXML(buf.Bytes())
	Expect(err).ToNot(HaveOccurred())

	return doc
}

func getTable(control string) xml.Node {
	doc := docFixture(control)
	tables, err := doc.Search("//w:tbl")
	Expect(err).NotTo(HaveOccurred())
	return tables[0]
}

var _ = Describe("SummaryTable", func() {
	Describe("Fill", func() {
		It("fills in the Responsible Role for controls", func() {
			table := getTable("AC-2")
			ct := SummaryTable{Root: table}
			openControlData := fixtures.LoadOpenControlFixture()

			ct.Fill(openControlData)

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})

		It("fills in the Responsible Role for control enhancements", func() {
			table := getTable("AC-2 (1)")
			ct := SummaryTable{Root: table}
			openControlData := fixtures.LoadOpenControlFixture()

			ct.Fill(openControlData)

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})
	})
})
