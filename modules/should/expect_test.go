package should_test

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/clearblade/cbtest/mocks"
	"github.com/clearblade/cbtest/modules/should"
	"github.com/clearblade/cbtest/modules/should/to"
)

func TestExpect(t *testing.T) {

	s := []string{"foo", "bar", "baz"}

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Errorf", mock.Anything, mock.Anything).Return()

	assert.True(t, should.ExpectE(mockT, s, to.ConsistOf("baz", "bar", "foo"))) // ordering doesn't matter
}

func TestExpect_WithGomega(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	assert.True(t, should.ExpectE(mockT, 10, gomega.BeNumerically(">", 5)))
	assert.True(t, should.ExpectE(mockT, 10, gomega.BeNumerically("==", 10)))
	assert.False(t, should.ExpectE(mockT, 10, gomega.BeNumerically(">", 15)))
}
