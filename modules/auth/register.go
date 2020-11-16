package auth

import (
	"strings"

	"github.com/stretchr/testify/require"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
)

// RegisterDev registers the given developer into the platform given by the config.
// Panics on failure.
func RegisterDev(t cbtest.T, provider config.Provider, email, password string) {
	t.Helper()
	err := cbRegisterDeveloper(t, provider, email, password)
	require.NoError(t, err)
}

// RegisterDevE registers the given developer into the platform given by the config.
// Returns error on failure.
func RegisterDevE(t cbtest.T, provider config.Provider, email, password string) error {
	t.Helper()
	err := cbRegisterDeveloper(t, provider, email, password)
	return err
}

// cbRegisterDeveloper registers a new developer in the system if it doesn't
// exists already.
func cbRegisterDeveloper(t cbtest.T, provider config.Provider, email, password string) error {
	t.Helper()

	config := provider.Provide()
	devClient := cb.NewDevClientWithAddrs(config.PlatformURL, config.MessagingURL, "", "")

	firstname := "cbtest"
	lastname := "cbtest"
	org := "ClearBlade, Inc."
	regkey := config.RegistrationKey

	_, err := devClient.RegisterDevUserWithKey(email, password, firstname, lastname, org, regkey)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}

	return nil
}

// RegisterUser registers the given user into the system given by the config.
// Panics on failure.
func RegisterUser(t cbtest.T, provider config.Provider, email, password string) {
	t.Helper()
	err := cbRegisterUser(t, provider, email, password)
	require.NoError(t, err)
}

// RegisterUserE registers the given user into the system given by the config.
// Returns error on failure.
func RegisterUserE(t cbtest.T, provider config.Provider, email, password string) error {
	t.Helper()
	err := cbRegisterUser(t, provider, email, password)
	return err
}

// cbRegisterUser registers a new user in the system if it doesn't exists already.
func cbRegisterUser(t cbtest.T, provider config.Provider, email, password string) error {
	t.Helper()

	devClient, err := LoginAsDevE(t, provider)
	if err != nil {
		return err
	}

	config := provider.Provide()

	_, err = devClient.RegisterNewUser(email, password, config.SystemKey, config.SystemSecret)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}

	user, err := devClient.GetUserInfo(config.SystemKey, email)
	if err != nil {
		return err
	}

	err = devClient.AddUserToRoles(config.SystemKey, user["user_id"].(string), []string{"Administrator"})
	if err != nil {
		return err
	}

	return nil
}

// RegisterDevice registers the given device into the system.
// Panics on failure.
func RegisterDevice(t cbtest.T, provider config.Provider, name, activeKey string) {
	t.Helper()
	err := RegisterDeviceE(t, provider, name, activeKey)
	require.NoError(t, err)
}

// RegisterDeviceE registers the given device into the system.
// Returns error on failure.
func RegisterDeviceE(t cbtest.T, provider config.Provider, name, activeKey string) error {
	t.Helper()
	err := cbRegisterDevice(t, provider, name, activeKey)
	return err
}

// cbRegisterDevice registers a new device in the system if it doesn't exists already.
func cbRegisterDevice(t cbtest.T, provider config.Provider, name, activeKey string) error {
	t.Helper()

	devClient, err := LoginAsDevE(t, provider)
	if err != nil {
		return err
	}

	data := cbMakeRegisterDeviceData(name, activeKey)
	config := provider.Provide()

	_, err = devClient.CreateDevice(config.SystemKey, name, data)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}

	err = devClient.AddDeviceToRoles(config.SystemKey, name, []string{"Administrator"})
	if err != nil {
		return err
	}

	return nil
}

// cbMakeRegisterDeviceData creates the payload required to create a new device.
// For some reason we need to specify these this when creting a device.
func cbMakeRegisterDeviceData(name, activeKey string) map[string]interface{} {
	return map[string]interface{}{
		"active_key":             activeKey,
		"allow_certificate_auth": true,
		"allow_key_auth":         true,
		"enabled":                true,
		"name":                   name,
		"type":                   "",
		"state":                  "",
	}
}
