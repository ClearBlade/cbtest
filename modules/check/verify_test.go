package check_test

import (
	"fmt"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/clearblade/cbtest/mocks"
	"github.com/clearblade/cbtest/modules/check"
)

var testingT = &mocks.T{}

func init() {
	testingT.On("Helper").Return()
	testingT.On("Errorf", mock.Anything, mock.Anything).Return()
}

func ExampleVerify_withGomega() {
	fmt.Println(check.VerifyE(testingT, 10, gomega.BeNumerically(">", 5)))
	fmt.Println(check.VerifyE(testingT, 10, gomega.BeNumerically("==", 10)))
	fmt.Println(check.VerifyE(testingT, 10, gomega.BeNumerically(">", 15)))
	// Output:
	// true
	// true
	// false
}
