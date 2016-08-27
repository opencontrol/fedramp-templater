package templater_test

import (
	"bytes"
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/fixtures"
	"github.com/opencontrol/fedramp-templater/ssp"
	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/reporter"
)

func extractDiffReport(reporters []reporter.Reporter) string {
	report := &bytes.Buffer{}
	for _, rept := range reporters {
		rept.WriteTextTo(report)
	}
	return report.String()
}

var _ = Describe("Templater", func() {
	Describe("TemplatizeSSP", func() {
		It("fills in the Responsible Role fields", func() {
			doc := fixtures.LoadSSP("FedRAMP_ac-2-1_v2.1.docx")
			defer doc.Close()
			openControlData := fixtures.LoadOpenControlFixture()

			err := TemplatizeSSP(doc, openControlData)

			Expect(err).NotTo(HaveOccurred())
			content := doc.Content()
			Expect(content).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})
	})

	Describe("DiffSSP", func() {
		It("should warn the user if the current SSP contains a responsible role that conflicts with the "+
			"responsbile role in the YAML", func() {

			By("Loading the SSP with the Responsible Role being 'OpenControl Role Placeholder' " +
				"for Control 'AC-2 (1)'")
			sspPath := filepath.Join("..", "fixtures", "FedRAMP_ac-2-1_v2.1.docx")
			s, err := ssp.Load(sspPath)
			Expect(err).NotTo(HaveOccurred())
			defer s.Close()

			By("Loading the data from the opencontrol workspace with the Responsible Role being " +
				"'Amazon Elastic Compute Cloud: AWS Staff' for Control 'AC-2 (1)'")
			openControlData := fixtures.LoadOpenControlFixture()

			By("Calling 'diff' on the SSP")
			diffInfo, err := DiffSSP(s, openControlData)
			Expect(err).NotTo(HaveOccurred())

			By("extracting the report, it should find the difference in responsible " +
				"roles and return it.")
			report := extractDiffReport(diffInfo)
			Expect(report).To(Equal("Control: AC-2 (1). " +
				"Responsible Role in SSP: \"OpenControl Role Placeholder\". " +
				"Responsible Role in YAML: \"Amazon Elastic Compute Cloud: AWS Staff\".\n"))
		})
	})
})
