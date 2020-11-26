package check

import (
	"fmt"

	"github.com/clearblade/cbtest/mocks"
)

var testingT = &mocks.T{}

func init() {
	testingT.On("Helper").Return()
}

func ExampleAnything() {
	fmt.Println(VerifyE(testingT, 0, Anything()))
	fmt.Println(VerifyE(testingT, "foo", Anything()))
	// Output:
	// true
	// true
}
