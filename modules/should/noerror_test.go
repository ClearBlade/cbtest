package should

import (
	"fmt"
	"testing"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNoErrorESucceeds(t *testing.T) {

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Errorf", mock.Anything, mock.Anything, mock.Anything).Return()

	assert.True(t, NoErrorE(mockT, nil))
	assert.False(t, NoErrorE(mockT, fmt.Errorf("some error")))
}
