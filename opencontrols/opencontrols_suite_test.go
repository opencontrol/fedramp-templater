package opencontrols_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOpencontrols(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opencontrols Suite")
}
