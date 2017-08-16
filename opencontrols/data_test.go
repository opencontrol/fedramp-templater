package opencontrols_test

import (
	"github.com/opencontrol/fedramp-templater/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Data", func() {
	Describe("GetNarrative", func() {
		It("returns the relevant singular narrative", func() {
			data := fixtures.LoadOpenControlFixture()
			result := data.GetNarrative("AC-2 (1)", "")
			Expect(result).To(Equal("Amazon Elastic Compute Cloud\nJustification in narrative form for AC-2 (1)\n"))
		})

		It("returns the relevant narrative section", func() {
			data := fixtures.LoadOpenControlFixture()
			result := data.GetNarrative("AC-2", "a")
			Expect(result).To(Equal("Amazon Elastic Compute Cloud\nJustification in narrative form A for AC-2\n"))
		})

		It("returns the correctly-merged multiple lines", func() {
			data := fixtures.LoadOpenControlFixture()
			result := data.GetNarrative("AC-2 (2)", "")
			Expect(result).To(MatchRegexp("(?m)\nNEWLINE Test 1.*? NO_NEWLINE Test 2.*?\nNEWLINE Test 3.*? NO_NEWLINE Test 4"))
		})
	})
})
