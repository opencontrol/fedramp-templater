package templater_test

import (
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/ssp"
	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Templater", func() {
	Describe("TemplatizeSsp", func() {
		It("fills in the Responsible Role fields", func() {
			path := filepath.Join("..", "fixtures", "FedRAMP_ac-2_v2.1.docx")
			s, err := ssp.Load(path)
			Expect(err).NotTo(HaveOccurred())
			defer s.Close()

			err = TemplatizeSsp(s)

			Expect(err).NotTo(HaveOccurred())
			content := s.Content()
			Expect(content).To(ContainSubstring(`Responsible Role: {{getResponsibleRole "NIST-800-53" "AC-2"}}`))
			Expect(content).To(ContainSubstring(`Responsible Role: {{getResponsibleRole "NIST-800-53" "AC-2 (1)"}}`))
		})
	})
})
