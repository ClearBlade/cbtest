package to_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/clearblade/cbtest/mocks"
	"github.com/clearblade/cbtest/modules/should"
	"github.com/clearblade/cbtest/modules/should/to"
)

func TestAnything(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()

	assert.True(t, should.ExpectE(mockT, nil, to.Anything()))
	assert.True(t, should.ExpectE(mockT, 0, to.Anything()))
	assert.True(t, should.ExpectE(mockT, "foo", to.Anything()))
}
