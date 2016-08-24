package templater_test

import (
	"path/filepath"

	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/ssp"
	. "github.com/opencontrol/fedramp-templater/templater"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Templater", func() {
	Describe("TemplatizeSSP", func() {
		It("fills in the Responsible Role fields", func() {
			sspPath := filepath.Join("..", "fixtures", "FedRAMP_ac-2-1_v2.1.docx")
			s, err := ssp.Load(sspPath)
			Expect(err).NotTo(HaveOccurred())
			defer s.Close()

			openControlDir := filepath.Join("..", "fixtures", "opencontrols")
			openControlDir, err = filepath.Abs(openControlDir)
			Expect(err).NotTo(HaveOccurred())
			openControlData, errors := opencontrols.LoadFrom(openControlDir)
			for _, err := range errors {
				Expect(err).NotTo(HaveOccurred())
			}

			err = TemplatizeSSP(s, openControlData)

			Expect(err).NotTo(HaveOccurred())
			content := s.Content()
			Expect(content).To(ContainSubstring(`Responsible Role: Amazon Elastic Compute Cloud: AWS Staff`))
		})
	})
	Describe("DiffSSP", func() {
		It("should warn the user if the current SSP contains a responsible role that conflicts with the " +
			"responsbile role in the YAML", func() {

			By("Loading the SSP with the Responsible Role being 'OpenControl Role Placeholder' " +
				"for Control 'AC-2 (1)'")
			sspPath := filepath.Join("..", "fixtures", "FedRAMP_ac-2-1_v2.1.docx")
			s, err := ssp.Load(sspPath)
			Expect(err).NotTo(HaveOccurred())
			defer s.Close()

			By("Loading the data from the opencontrol workspace with the Responsible Role being " +
				"'Amazon Elastic Compute Cloud: AWS Staff' for Control 'AC-2 (1)'")
			openControlDir := filepath.Join("..", "fixtures", "opencontrols")
			openControlDir, err = filepath.Abs(openControlDir)
			Expect(err).NotTo(HaveOccurred())
			openControlData, errors := opencontrols.LoadFrom(openControlDir)
			for _, err := range errors {
				Expect(err).NotTo(HaveOccurred())
			}

			By("Calling 'diff' on the SSP, it should find the difference in responsible " +
				"roles and return it.")
			diffInfo, err := DiffSSP(s, openControlData)
			Expect(diffInfo).To(ContainElement("Control: AC-2 (1). " +
				"Responsible Role in doc :\"OpenControl Role Placeholder\". " +
				"Responsible Role in YAML \"Amazon Elastic Compute Cloud: AWS Staff\n\"\n"))
		})
	})
})
