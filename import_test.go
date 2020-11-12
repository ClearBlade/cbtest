package cbtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSystemValidSystemSucceeds(t *testing.T) {
	err := checkSystem("golden/sys")
	assert.NoError(t, err)
}

func TestCheckSystemNonSystemFails(t *testing.T) {
	err := checkSystem("golden/foo")
	assert.Error(t, err)
}
