package lookup

import (
	"fmt"
	"testing"

	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testingT = &mocks.T{}

func init() {
	testingT.On("Helper").Return()
}

func ExamplePath() {
	m := map[string]interface{}{
		"nested": map[string]interface{}{
			"value": "foo",
		},
	}

	value := Path(testingT, m, "nested.value")
	fmt.Println(value)
	// Output:
	// foo
}

func ExamplePathI() {
	m := map[string]interface{}{
		"nested": map[string]interface{}{
			"value": "bar",
		},
	}

	value := PathI(testingT, m, "NESted.VALue")
	fmt.Println(value)
	// Output:
	// bar
}

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
