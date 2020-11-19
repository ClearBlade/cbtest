package ginkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoTest Suite")
}
