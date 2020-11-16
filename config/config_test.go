package config

import (
	"reflect"
	"strings"
	"testing"

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
		// NOTE: add more flags to the test above this line
	}

	for _, tt := range tests {

		reflect.ValueOf(tt.flag).Elem().Set(reflect.ValueOf(tt.want))
		config, err := ObtainConfig()
		require.NoError(t, err)

		value, err := lookup.LookupString(config, tt.cpath)
		require.NoError(t, err)

		assert.Equal(t, tt.want, value.Interface())
	}
}
