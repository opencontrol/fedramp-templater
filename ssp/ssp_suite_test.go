package ssp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSSP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SSP Suite")
}
