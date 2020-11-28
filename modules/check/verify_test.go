package check_test

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/clearblade/cbtest/mocks"
	"github.com/clearblade/cbtest/modules/check"
)

func TestVerify(t *testing.T) {

	s := []string{"foo", "bar", "baz"}

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Errorf", mock.Anything, mock.Anything).Return()

	assert.True(t, check.VerifyE(mockT, s, check.ConsistOf("baz", "bar", "foo"))) // ordering doesn't matter
}

func TestVerify_WithGomega(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	assert.True(t, check.VerifyE(mockT, 10, gomega.BeNumerically(">", 5)))
	assert.True(t, check.VerifyE(mockT, 10, gomega.BeNumerically("==", 10)))
	assert.False(t, check.VerifyE(mockT, 10, gomega.BeNumerically(">", 15)))
}