package check

import (
	"testing"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAnything(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()

	assert.True(t, ExpectE(mockT, nil, Anything()))
	assert.True(t, ExpectE(mockT, 0, Anything()))
	assert.True(t, ExpectE(mockT, "foo", Anything()))
}
