package templater_test

import (
	"github.com/opencontrol/fedramp-templater/fixtures"
	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
})
