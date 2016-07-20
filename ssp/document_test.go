package ssp_test

import (
	"path/filepath"

	. "github.com/opencontrol/fedramp-templater/ssp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SSP", func() {
	Describe("Load", func() {
		It("gets the content from the doc", func() {
			path := filepath.Join("..", "fixtures", "FedRAMP_ac-2-1_v2.1.docx")
			s, err := Load(path)
			Expect(err).NotTo(HaveOccurred())
			defer s.Close()

			Expect(s.Content()).To(ContainSubstring("Control Enhancement"))
		})

		It("give an error when the doc isn't found", func() {
			_, err := Load("non-existent.docx")
			Expect(err).To(HaveOccurred())
		})
	})
})
