// Package ginkgo showcases how cbtest can be used with alternative test
// frameworks thanks to the usage of the cbtest.T interface.
// See: https://github.com/onsi/ginkgo
package ginkgo

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/service"
	"github.com/clearblade/cbtest/modules/system"
)

const (
	AdderService = "adder"
)

func T() GinkgoTInterface {
	return GinkgoT()
}

var _ = Describe("Adder test", func() {

	table := []struct {
		lhs  float64
		rhs  float64
		want float64
	}{
		{0, 0, 0},
		{1, 1, 2},
		{3, 4, 7},
	}

	// references that we use in our suite
	var s *system.EphemeralSystem
	var devClient *cb.DevClient

	BeforeSuite(func() {

		// import into new system
		s = system.UseOrImport(T(), "./extra")

		// obtain developer client from the ephemeral system
		devClient = auth.LoginAsDev(T(), s)

	})

	AfterSuite(func() {

		// destroy the system after the test
		system.Destroy(T(), s)

	})

	for _, tt := range table {

		// NOTE: capture tt
		// see: https://www.calhoun.io/gotchas-and-common-mistakes-with-closures-in-go/#variables-declared-in-for-loops-are-passed-by-reference
		tt := tt
		name := fmt.Sprintf("%.2f + %.2f = %.2f", tt.lhs, tt.rhs, tt.want)

		It(name, func() {
			adderCase(s, devClient, tt.lhs, tt.rhs, tt.want)
		})

	}
})

func adderCase(s *system.EphemeralSystem, devClient *cb.DevClient, lhs, rhs, want float64) {

	// payload that we will send to the adder service
	payload := map[string]interface{}{"lhs": lhs, "rhs": rhs}

	// call the service
	resp, err := devClient.CallService(s.SystemKey(), AdderService, payload, false)
	Expect(err).To(BeNil())

	// assert response from service
	service.AssertResponseEqual(T(), want, resp)
}
