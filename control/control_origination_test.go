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
			// Check error
			Expect(err).ToNot(HaveOccurred())
			// Check number of control origination.
			Expect(len(co.origins)).To(Equal(7))

			// Find the checked service provided corporate origination.
			Expect(co.origins[0].getTextValue()).To(ContainSubstring(serviceProviderCorporateOrigination))
			Expect(co.origins[0].isChecked()).To(Equal(true))

			// Find the unchecked service provided system specific origination.
			Expect(co.origins[1].getTextValue()).
				To(ContainSubstring(serviceProviderSystemSpecificOrigination))
			Expect(co.origins[1].isChecked()).To(Equal(false))

			// Find the unchecked service provided hybrid origination.
			Expect(co.origins[2].getTextValue()).To(ContainSubstring(serviceProviderHybridOrigination))
			Expect(co.origins[2].isChecked()).To(Equal(false))

			// Find the unchecked configured by customer origination.
			Expect(co.origins[3].getTextValue()).To(ContainSubstring(configuredByCustomerOrigination))
			Expect(co.origins[3].isChecked()).To(Equal(false))

			// Find the unchecked provided by customer origination.
			Expect(co.origins[4].getTextValue()).To(ContainSubstring(providedByCustomerOrigination))
			Expect(co.origins[4].isChecked()).To(Equal(false))

			// Find the unchecked shared origination.
			Expect(co.origins[5].getTextValue()).To(ContainSubstring(sharedOrigination))
			Expect(co.origins[5].isChecked()).To(Equal(false))

			// Find the unchecked inherited origination.
			Expect(co.origins[6].getTextValue()).To(ContainSubstring(inheritedOrigination))
			Expect(co.origins[6].isChecked()).To(Equal(false))
		})
	})
})
