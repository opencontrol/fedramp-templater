package control

import (
	"bytes"
	"text/template"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/docx/helper"
	"github.com/opencontrol/fedramp-templater/ssp"
	"github.com/opencontrol/fedramp-templater/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/reporter"
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
	// replicate what ssp.Document's SummaryTables() method is doing, except that this source isn't a full Word doc
	tables, err := doc.Search(ssp.SummaryTablesXPath)
	Expect(err).NotTo(HaveOccurred())
	return tables[0]
}

var _ = Describe("SummaryTable", func() {
	Describe("Fill", func() {
		It("fills in the Responsible Role for controls", func() {
			table := getTable("AC-2")
			st := NewSummaryTable(table)
			openControlData := fixtures.LoadOpenControlFixture()

			st.Fill(openControlData)

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})

		It("fills in the Responsible Role for control enhancements", func() {
			table := getTable("AC-2 (1)")
			st := NewSummaryTable(table)
			openControlData := fixtures.LoadOpenControlFixture()

			st.Fill(openControlData)

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})
		It("fills in the control origination", func() {
			table := getTable("AC-2")
			st := NewSummaryTable(table)
			openControlData := fixtures.LoadOpenControlFixture()

			st.Fill(openControlData)

			origination, err := newControlOrigination(st)
			Expect(err).ToNot(HaveOccurred())
			Expect(origination.origins[sharedOrigination].IsChecked()).To(Equal(true))
		})
	})

	Describe("Diff", func() {
		It("detects no diff when the value of responsible role is empty", func() {
			table := getTable("AC-2")
			st := NewSummaryTable(table)
			openControlData := fixtures.LoadOpenControlFixture()
			diff, err := st.Diff(openControlData)

			Expect(diff).To(Equal([]reporter.Reporter{}))
			Expect(err).To(BeNil())
		})
	})
})
