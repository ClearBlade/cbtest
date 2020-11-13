package cbtest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfigCamelCaseFieldSucceeds(t *testing.T) {

	rawConfig := `{ "platformUrl": "foo.bar.baz" }`
	config, err := ReadConfig(strings.NewReader(rawConfig))
	require.NoError(t, err)

	assert.Equal(t, "foo.bar.baz", config.PlatformURL)
}

func TestReadConfigSnakeCaseFieldFails(t *testing.T) {

	rawConfig := `{ "platform_url": "foo.bar.baz" }`
	config, err := ReadConfig(strings.NewReader(rawConfig))
	require.NoError(t, err)

	assert.NotEqual(t, "foo.bar.baz", config.PlatformURL)
}
