package lookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathE_ValidPathSucceeds(t *testing.T) {

	m := map[string]interface{}{
		"nested": map[string]interface{}{
			"value": "foo",
		},
	}

	value, err := PathE(t, m, "nested.value")
	require.NoError(t, err)
	assert.Equal(t, "foo", value)
}

func TestPathE_InvalidPathFails(t *testing.T) {

	m := map[string]interface{}{
		"nested": map[string]interface{}{
			"value": "foo",
		},
	}

	_, err := PathE(t, m, "nested.unknown")
	require.Error(t, err)
}

func TestPathIE_ValidPathSucceeds(t *testing.T) {

	m := map[string]interface{}{
		"nested": map[string]interface{}{
			"value": "foo",
		},
	}

	value, err := PathIE(t, m, "NESted.VALue")
	require.NoError(t, err)
	assert.Equal(t, "foo", value)
}

func TestPathIE_InvalidPathFails(t *testing.T) {

	m := map[string]interface{}{
		"nested": map[string]interface{}{
			"value": "foo",
		},
	}

	_, err := PathIE(t, m, "NESted.UNKnown")
	require.Error(t, err)
}
