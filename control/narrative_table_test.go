package control_test

import (
	. "github.com/opencontrol/fedramp-templater/control"
	"github.com/opencontrol/fedramp-templater/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NarrativeTable", func() {
	Describe("SectionRows", func() {
		It("returns the correct number for a singular narrative", func() {
			doc := fixtures.LoadSSP("FedRAMP_ac-2-1_v2.1.docx")
			defer doc.Close()
			root, err := doc.NarrativeTable("AC-2 (1)")
			Expect(err).NotTo(HaveOccurred())

			table := NewNarrativeTable(root)
			sections, err := table.SectionRows()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(sections)).To(Equal(1))
		})

		It("returns the correct number for multiple narrative sections", func() {
			doc := fixtures.LoadSSP("FedRAMP_ac-2_v2.1.docx")
			defer doc.Close()
			root, err := doc.NarrativeTable("AC-2")
			Expect(err).NotTo(HaveOccurred())

			table := NewNarrativeTable(root)
			sections, err := table.SectionRows()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(sections)).To(Equal(11))
		})
	})
})
