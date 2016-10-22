package control

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/fixtures"
)

var _ = Describe("controlOrignation", func() {
	Describe("newControlOrignation", func() {
		It("should be called when calling NewSummaryTable and should have the origins", func() {
			doc := fixtures.LoadSSP("FedRAMP_ac-2-1_v2.1.docx")
			defer doc.Close()
			tables, err := doc.SummaryTables()
			Expect(err).NotTo(HaveOccurred())
			st, err := NewSummaryTable(tables[0])
			Expect(err).NotTo(HaveOccurred())
			// Check number of control origination.
			Expect(len(st.originTable.origins)).To(Equal(7))
		})
	})
})
