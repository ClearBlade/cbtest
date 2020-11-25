package config

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/clearblade/cbtest/internal/fsutil"
	"github.com/mcuadros/go-lookup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUseOrDefaultSucceeds(t *testing.T) {
	assert.Equal(t, "foo", useOrDefault("foo", "bar"))
	assert.Equal(t, "bar", useOrDefault("", "bar"))
}

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

func TestObtainConfigOverrides(t *testing.T) {
	tests := []struct {
		flag  interface{}
		cpath string
		want  interface{}
	}{
		{flagPlatformURL, "PlatformURL", "platform-url-override"},
		{flagMessagingURL, "MessagingURL", "messaging-url-override"},
		{flagRegistrationKey, "RegistrationKey", "messaging-key-override"},
		{flagSystemKey, "SystemKey", "system-key-override"},
		{flagSystemSecret, "SystemSecret", "system-secret-override"},
		{flagDevEmail, "Developer.Email", "dev-email-override"},
		{flagDevPassword, "Developer.Password", "dev-password-override"},
		{flagUserEmail, "User.Email", "user-email-override"},
		{flagUserPassword, "User.Password", "user-password-override"},
		{flagDeviceName, "Device.Name", "device-name-override"},
		{flagDeviceActiveKey, "Device.ActiveKey", "device-active-key-override"},
		{flagImportUsers, "Import.ImportUsers", false},
		{flagImportRows, "Import.ImportRows", false},
		{flagConfigOut, "Out", "config-out-override"},
		// NOTE: add more flags to the test above this line
	}

	for _, tt := range tests {

		// saves old flag value and sets new one
		oldFlagValue := reflect.ValueOf(tt.flag).Elem().Interface()
		reflect.ValueOf(tt.flag).Elem().Set(reflect.ValueOf(tt.want))

		// obtain config (flag should override whatever we get)
		config, err := ObtainConfig(t)
		require.NoError(t, err)

		// obtain the overridden value from the obtained config
		value, err := lookup.LookupString(config, tt.cpath)
		require.NoError(t, err)

		// make sure flag and value in config are equal
		assert.Equal(t, tt.want, value.Interface())

		// recover old flag value
		reflect.ValueOf(tt.flag).Elem().Set(reflect.ValueOf(oldFlagValue))
	}
}

func TestOutFieldShouldSaveSucceeds(t *testing.T) {
	config := GetDefaultConfig()
	assert.False(t, config.ShouldSave())
	config.Out = "some-path"
	assert.True(t, config.ShouldSave())
}

func TestSaveConfig_SpecifyingOutSucceeds(t *testing.T) {

	tempdir, cleanup := fsutil.MakeTempDir()
	defer cleanup()

	config := GetDefaultConfig()
	config.Out = filepath.Join(tempdir, "config.json")

	assert.False(t, fsutil.IsFile(config.Out))
	SaveConfig(t, config)
	assert.True(t, fsutil.IsFile(config.Out))
}

func TestSaveConfig_generateOutputName(t *testing.T) {
	outputName := generateOutputName(t)
	assert.NotEmpty(t, outputName)
}
