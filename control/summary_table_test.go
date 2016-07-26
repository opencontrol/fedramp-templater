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
			doc := docFixture("AC-2")
			tables, _ := doc.Search("//w:tbl")
			table := tables[0]

			ct := SummaryTable{Root: table}
			openControlData := openControlFixture()
			ct.Fill(openControlData)

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})

		It("fills in the Responsible Role for control enhancements", func() {
			doc := docFixture("AC-2 (1)")
			tables, _ := doc.Search("//w:tbl")
			table := tables[0]

			ct := SummaryTable{Root: table}
			openControlData := openControlFixture()
			ct.Fill(openControlData)

			Expect(table.Content()).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})
	})
	Describe("Diff", func() {
		It("detects no diff when the value of responsible role is empty", func() {
			doc := docFixture("AC-2")
			tables, _ := doc.Search("//w:tbl")
			table := tables[0]

			ct := SummaryTable{Root: table}
			openControlData := openControlFixture()
			diff, err := ct.Diff(openControlData)

			Expect(diff).To(Equal([]reporter.Reporter{}))
			Expect(err).To(BeNil())
		})
	})
})
