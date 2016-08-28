package control_test

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"
	// using fork because of https://github.com/moovweb/gokogiri/pull/93#issuecomment-215582446
	"github.com/jbowtie/gokogiri/xml"

	. "github.com/opencontrol/fedramp-templater/control"
	"github.com/opencontrol/fedramp-templater/docx/helper"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/reporter"
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

func openControlFixturePath() string {
	path := filepath.Join("..", "fixtures", "opencontrols")
	path, err := filepath.Abs(path)
	Expect(err).NotTo(HaveOccurred())
	_, err = os.Stat(path)
	Expect(err).NotTo(HaveOccurred())

	return path
}

func openControlFixture() opencontrols.Data {
	path := openControlFixturePath()
	data, errors := opencontrols.LoadFrom(path)
	for _, err := range errors {
		Expect(err).NotTo(HaveOccurred())
	}

	return data
}

var _ = Describe("SummaryTable", func() {
	Describe("Fill", func() {
		It("fills in the Responsible Role for controls", func() {
			table := getTable("AC-2")
			st := SummaryTable{Root: table}
			openControlData := openControlFixture()

			st.Fill(openControlData)

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})

		It("fills in the Responsible Role for control enhancements", func() {
			table := getTable("AC-2 (1)")
			st := SummaryTable{Root: table}
			openControlData := openControlFixture()

			st.Fill(openControlData)

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})
	})

	Describe("Diff", func() {
		It("detects no diff when the value of responsible role is empty", func() {
			table := getTable("AC-2")
			st := SummaryTable{Root: table}
			openControlData := openControlFixture()

			diff, err := st.Diff(openControlData)

			Expect(diff).To(Equal([]reporter.Reporter{}))
			Expect(err).To(BeNil())
		})
	})
})
