package control

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/reporter"
	"bytes"
)

func createFakeStdOut() *bytes.Buffer {
	return &bytes.Buffer{}
}

var _ = Describe("Reporter", func() {
	Describe("NewDiff", func() {
		It("should return a new reporter. (Will fail to compile if it doesn't comply with the interface)", func() {
			var diff reporter.Reporter
			diff = NewDiff("control", "myfield", "sspValue", "yamlValue")
			_ = diff
		})
	})
	Describe("WriteTextTo", func() {
		It("should write data in a plain text to the writer", func() {
			var diff reporter.Reporter
			diff = NewDiff("control", "myfield", "sspValue", "yamlValue")
			fakeConsole := createFakeStdOut()
			diff.WriteTextTo(fakeConsole)
			Expect(fakeConsole.String()).To(Equal("Control: control. myfield in SSP: \"sspValue\". myfield in YAML: \"yamlValue\".\n"))
		})
	})
})
