package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSystem_ValidSystemSucceeds(t *testing.T) {
	err := cbCheckSystem("./testdata/sys")
	assert.NoError(t, err)
}

func TestCheckSystem_BadSystemFails(t *testing.T) {
	err := cbCheckSystem("./testdata/nosys")
	assert.Error(t, err)
}
