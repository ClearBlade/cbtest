package check_test

import (
	"fmt"

	"github.com/onsi/gomega"

	"github.com/clearblade/cbtest/mocks"
	"github.com/clearblade/cbtest/modules/check"
)

var testingT = &mocks.T{}

func init() {
	testingT.On("Helper").Return()
}

func ExampleVerify_withGomega() {
	fmt.Println(check.VerifyE(testingT, 10, gomega.BeNumerically(">", 5)))
	// Output:
	// true
}
