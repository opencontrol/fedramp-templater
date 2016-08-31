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

			// Find the unchecked service provided corporate origination.
			Expect(co.origins[serviceProviderCorporateOrigination].GetTextValue()).
				To(ContainSubstring(serviceProviderCorporateOrigination))
			Expect(co.origins[serviceProviderCorporateOrigination].IsChecked()).To(Equal(false))

			// Find the unchecked service provided system specific origination.
			Expect(co.origins[serviceProviderSystemSpecificOrigination].GetTextValue()).
				To(ContainSubstring(serviceProviderSystemSpecificOrigination))
			Expect(co.origins[serviceProviderSystemSpecificOrigination].IsChecked()).To(Equal(false))

			// Find the unchecked service provided hybrid origination.
			Expect(co.origins[serviceProviderHybridOrigination].GetTextValue()).
				To(ContainSubstring(serviceProviderHybridOrigination))
			Expect(co.origins[serviceProviderHybridOrigination].IsChecked()).To(Equal(false))

			// Find the unchecked configured by customer origination.
			Expect(co.origins[configuredByCustomerOrigination].GetTextValue()).
				To(ContainSubstring(configuredByCustomerOrigination))
			Expect(co.origins[configuredByCustomerOrigination].IsChecked()).To(Equal(false))

			// Find the unchecked provided by customer origination.
			Expect(co.origins[providedByCustomerOrigination].GetTextValue()).
				To(ContainSubstring(providedByCustomerOrigination))
			Expect(co.origins[providedByCustomerOrigination].IsChecked()).To(Equal(false))

			// Find the unchecked shared origination.
			Expect(co.origins[sharedOrigination].GetTextValue()).
				To(ContainSubstring(sharedOrigination))
			Expect(co.origins[sharedOrigination].IsChecked()).To(Equal(false))

			// Find the unchecked inherited origination.
			Expect(co.origins[inheritedOrigination].GetTextValue()).
				To(ContainSubstring(inheritedOrigination))
			Expect(co.origins[inheritedOrigination].IsChecked()).To(Equal(false))
		})
	})
})
