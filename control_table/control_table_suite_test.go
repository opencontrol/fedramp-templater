package control_table_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestControlTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ControlTable Suite")
}
