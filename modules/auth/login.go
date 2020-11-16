package auth

import (
	"fmt"

	cb "github.com/clearblade/Go-SDK"
	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/config"
)

// LoginAsDev logs into the system as a Developer (given by config).
// Panics on failure.
func LoginAsDev(t cbtest.T, provider config.Provider) *cb.DevClient {
	t.Helper()
	devClient, err := LoginAsDevE(t, provider)
	require.NoError(t, err)
	return devClient
}

// LoginAsDevE logs into the System as a Developer (given by config).
// Returns error on failure.
func LoginAsDevE(t cbtest.T, provider config.Provider) (*cb.DevClient, error) {
	t.Helper()
	config := provider.Provide()
	return LoginDevE(t, provider, config.Developer.Email, config.Developer.Password)
}

// LoginDev logs into the platform as a developer.
// Panics on failure.
func LoginDev(t cbtest.T, provider config.Provider, email, password string) *cb.DevClient {
	t.Helper()
	devClient, err := LoginDevE(t, provider, email, password)
	require.NoError(t, err)
	return devClient
}

// LoginDevE logs into the platform as a developer.
// Returns error on failure.
func LoginDevE(t cbtest.T, provider config.Provider, email, password string) (*cb.DevClient, error) {
	t.Helper()
	return cbLoginDev(provider, email, password)
}

func cbLoginDev(provider config.Provider, email, password string) (*cb.DevClient, error) {

	var err error
	config := provider.Provide()

	if !config.HasDeveloper() {
		return nil, fmt.Errorf("config does not have developer information")
	}

	devClient := cb.NewDevClientWithAddrs(config.PlatformURL, config.MessagingURL, email, password)
	_, err = devClient.Authenticate()
	if err != nil {
		return nil, err
	}

	return devClient, nil
}

// LoginAsUser logs into the system as a User (given by config).
// Panics on failure.
func LoginAsUser(t cbtest.T, provider config.Provider) *cb.UserClient {
	t.Helper()
	userClient, err := LoginAsUserE(t, provider)
	require.NoError(t, err)
	return userClient
}

// LoginAsUserE logs into the system as a User (given by config).
// Returns error on failure.
func LoginAsUserE(t cbtest.T, provider config.Provider) (*cb.UserClient, error) {
	t.Helper()
	config := provider.Provide()
	return LoginUserE(t, provider, config.User.Email, config.User.Password)
}

// LoginUser logs into the system as an User.
// Panics on error.
func LoginUser(t cbtest.T, provider config.Provider, email, password string) *cb.UserClient {
	t.Helper()
	userClient, err := LoginUserE(t, provider, email, password)
	require.NoError(t, err)
	return userClient
}

// LoginUserE logs into the system as an User.
// Returns error on failure.
func LoginUserE(t cbtest.T, provider config.Provider, email, password string) (*cb.UserClient, error) {
	t.Helper()
	return cbLoginUser(provider, email, password)
}

func cbLoginUser(provider config.Provider, email, password string) (*cb.UserClient, error) {

	config := provider.Provide()

	if !config.HasUser() {
		return nil, fmt.Errorf("config does not have user information")
	}

	userClient := cb.NewUserClientWithAddrs(config.PlatformURL, config.MessagingURL, config.SystemKey, config.SystemSecret, email, password)
	_, err := userClient.Authenticate()
	if err != nil {
		return nil, err
	}

	return userClient, nil

}

// LoginAsDevice logs into the system as a Device (given by config).
// Panics on failure.
func LoginAsDevice(t cbtest.T, provider config.Provider) *cb.DeviceClient {
	t.Helper()
	deviceClient, err := LoginAsDeviceE(t, provider)
	require.NoError(t, err)
	return deviceClient
}

// LoginAsDeviceE logs into the system as a Device (given by config).
// Returns error on failure.
func LoginAsDeviceE(t cbtest.T, provider config.Provider) (*cb.DeviceClient, error) {
	t.Helper()
	config := provider.Provide()
	return LoginDeviceE(t, provider, config.Device.Name, config.Device.ActiveKey)
}

// LoginDevice logs into the system as an Device.
// Panics on error.
func LoginDevice(t cbtest.T, provider config.Provider, name, activeKey string) *cb.DeviceClient {
	t.Helper()
	deviceClient, err := LoginDeviceE(t, provider, name, activeKey)
	require.NoError(t, err)
	return deviceClient
}

// LoginDeviceE logs into the system as an Device.
// Returns error on failure.
func LoginDeviceE(t cbtest.T, provider config.Provider, name, activeKey string) (*cb.DeviceClient, error) {
	t.Helper()
	return cbLoginDevice(provider, name, activeKey)
}

func cbLoginDevice(provider config.Provider, name, activeKey string) (*cb.DeviceClient, error) {

	config := provider.Provide()

	if !config.HasDevice() {
		return nil, fmt.Errorf("config does not have device information")
	}

	deviceClient := cb.NewDeviceClientWithAddrs(config.PlatformURL, config.MessagingURL, config.SystemKey, config.SystemSecret, name, activeKey)
	_, err := deviceClient.Authenticate()
	if err != nil {
		return nil, err
	}

	return deviceClient, nil
}
