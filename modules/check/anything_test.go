package check

import (
	"testing"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAnything(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()

	assert.True(t, VerifyE(mockT, nil, Anything()))
	assert.True(t, VerifyE(mockT, 0, Anything()))
	assert.True(t, VerifyE(mockT, "foo", Anything()))
}
