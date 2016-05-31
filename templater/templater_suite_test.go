package templater_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTemplater(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Templater Suite")
}
