package npm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNPM_ValidNPMSucceeds(t *testing.T) {
	err := checkNPM("testdata/npm")
	assert.NoError(t, err)
}

func TestCheckNPM_NonValidNPMFails(t *testing.T) {
	err := checkNPM("testdata/nonpm")
	assert.Error(t, err)
}
