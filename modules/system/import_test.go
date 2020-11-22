package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSystem_ValidSystemSucceeds(t *testing.T) {
	err := checkSystem("testdata/sys")
	assert.NoError(t, err)
}

func TestCheckSystem_NonValidSystemFails(t *testing.T) {
	err := checkSystem("testdata/nosys")
	assert.Error(t, err)
}
