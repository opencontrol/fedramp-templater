package control


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/fixtures"
)

var _ = Describe("controlOrignation", func() {
	Describe("newControlOrignation", func() {
		It("should return", func() {
			doc := fixtures.LoadSSP("FedRAMP_ac-2-1_v2.1.docx")
			defer doc.Close()
			tables, err := doc.SummaryTables()
			Expect(err).NotTo(HaveOccurred())
			st := NewSummaryTable(tables[0])
			co, err := newControlOrigination(st)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(co.origins)).To(Equal(7))
			Expect(co.origins[0].getTextValue()).To(ContainSubstring(serviceProviderCorporate))
			Expect(co.origins[0].isChecked()).To(Equal(true))
			Expect(co.origins[1].getTextValue()).To(ContainSubstring(serviceProviderSystemSpecific))
			Expect(co.origins[1].isChecked()).To(Equal(false))
			Expect(co.origins[2].getTextValue()).To(ContainSubstring(serviceProviderHybrid))
			Expect(co.origins[2].isChecked()).To(Equal(false))
			Expect(co.origins[3].getTextValue()).To(ContainSubstring(configuredByCustomer))
			Expect(co.origins[3].isChecked()).To(Equal(false))
			Expect(co.origins[4].getTextValue()).To(ContainSubstring(providedByCustomer))
			Expect(co.origins[4].isChecked()).To(Equal(false))
			Expect(co.origins[5].getTextValue()).To(ContainSubstring(shared))
			Expect(co.origins[5].isChecked()).To(Equal(false))
			Expect(co.origins[6].getTextValue()).To(ContainSubstring(inherited))
			Expect(co.origins[6].isChecked()).To(Equal(false))
		})
	})
})
