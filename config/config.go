// Package config contains the structures, flags, and parsing of the cbtest
// configuration. Note that configuration is optional, and can be passed
// either using a config file or command line flags (with the latter taking
// precedence over the former).
//
// Configuration file
//
// A minimal configuration file looks like follows:
//
//     {
//       "platformUrl": "https://INSTANCE.clearblade.com",
//       "messagingUrl": "INSTANCE.clearblade.com:1883",
//       "registrationKey": "some-registration-key",
//       "systemKey": "abc123abc123abc123abc123abc123",
//       "systemSecret": "ABC123ABC123ABC123ABC123ABC123",
//       "developer": {
//         "email": "some-email@email.com",
//         "password": "some-password"
//       }
//     }
//
// Flags
//
// Since cbtest integrates with go test, flags can be specified as follows:
//
//     go test -v <PATH-TO-TEST> -args -cbtest.config <PATH-TO-CONFIG> [-cbtest.*]
//
package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/clearblade/cbtest"
)

// useOrDefault returns `value` if not empty, otherwise it returns `fallback`.
func useOrDefault(value, fallback string) string {
	if len(value) > 0 {
		return value
	}
	return fallback
}

// Config contains information about the system that is gonna be used for
// testing.
type Config struct {
	PlatformURL     string     `json:"platformUrl" mapstructure:"platformUrl"`
	MessagingURL    string     `json:"messagingUrl" mapstructure:"messagingUrl"`
	RegistrationKey string     `json:"registrationKey,omitempty" mapstructure:"registrationKey"`
	SystemKey       string     `json:"systemKey,omitempty" mapstructure:"systemKey"`
	SystemSecret    string     `json:"systemSecret,omitempty" mapstructure:"systemSecret"`
	Developer       *Developer `json:"developer,omitempty" mapstructure:"developer"`
	User            *User      `json:"user,omitempty" mapstructre:"user"`
	Device          *Device    `json:"device,omitempty" mapstructure:"device"`
	Import          *Import    `json:"import,omitempty" mapstructure:"import"`
	Out             string     `json:"-" mapstructure:"-"`
}

// Developer contains the developer credentials that must be provided if using
// an existing system.
type Developer struct {
	Email    string `json:"email" mapstructure:"email"`
	Password string `json:"password" mapstructure:"password"`
}

// User contains the credentials for an user in the system.
type User struct {
	Email    string `json:"email" mapstructure:"email"`
	Password string `json:"password" mapstructure:"password"`
}

// Device contains the credentials for a device in the system.
type Device struct {
	Name      string `json:"name" mapstructure:"name"`
	ActiveKey string `json:"activeKey" mapstructure:"activeKey"`
}

// Import contains import configuration values.
type Import struct {
	ImportUsers bool `json:"importUsers" mapstructure:"importUsers"`
	ImportRows  bool `json:"importRows" mapstructure:"importRows"`
}

// GetDefaultConfig returns a new *Config instance with default values.
func GetDefaultConfig() *Config {
	return &Config{
		PlatformURL:     "https://dev.clearblade.com",
		MessagingURL:    "dev.clearblade.com:1883",
		RegistrationKey: "",
		SystemKey:       "",
		SystemSecret:    "",
		Developer: &Developer{
			Email:    "cbtest@email.com",
			Password: "cbtestpassword",
		},
		User: &User{
			Email:    "cbtest@email.com",
			Password: "cbtestpassword",
		},
		Device: &Device{
			Name:      "cbtest-device",
			ActiveKey: "cbtestpassword",
		},
		Import: &Import{},
		Out:    "",
	}
}

// ReadConfigFromPath reads the config from the given path.
func ReadConfigFromPath(path string) (*Config, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ReadConfig(f)
}

// ReadConfig reads the config from the given reader.
func ReadConfig(r io.Reader) (*Config, error) {

	config := GetDefaultConfig()
	err := json.NewDecoder(r).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// WriteConfigToPath writes the config to the given path.
func WriteConfigToPath(path string, c *Config) error {

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return WriteConfig(f, c)
}

// WriteConfig writes the config to the given writer.
func WriteConfig(w io.Writer, c *Config) error {

	twoSpaces := "  "
	enc := json.NewEncoder(w)
	enc.SetIndent("", twoSpaces)

	err := enc.Encode(c)
	if err != nil {
		return err
	}

	return nil
}

// ObtainConfig returns the config that is gonna be used by cbtest. It uses
// either (1) default config, or (2) config read from the config path flag.
// Final values in the config are overridden by those provided in the flags.
func ObtainConfig(t cbtest.T) (*Config, error) {
	t.Helper()

	var config *Config
	var err error

	if HasConfig() {
		t.Logf("Reading config from path: %s", ConfigPath())
		config, err = ReadConfigFromPath(ConfigPath())
		if err != nil {
			return nil, fmt.Errorf("could not obtain config: %s", err)
		}
	} else {
		t.Logf("Using default config")
		config = GetDefaultConfig()
	}

	t.Logf("Overriding config from flags")
	config.overrideFromFlags()
	return config, nil
}

// SaveConfig saves the config to disk. It uses either (1) the config Out field
// if provided, or (2) a default path constructed from the given test object.
// Returns the output path and an error (if any).
func SaveConfig(t cbtest.T, c *Config) (string, error) {
	t.Helper()

	var err error

	outputPath := c.Out
	if strings.TrimSpace(outputPath) == "" {
		outputPath = generateOutputName(t)
	}

	t.Logf("Saving config to: %s", outputPath)
	err = WriteConfigToPath(outputPath, c)
	if err != nil {
		return "", fmt.Errorf("could not save config: %s", err)
	}

	return outputPath, nil
}

// generateOutputName generates an suitable output name for the config based
// on the test information
func generateOutputName(t cbtest.T) string {
	timestamp := time.Now().UTC().Unix()
	return fmt.Sprintf("cbtest-%s-%d.json", t.Name(), timestamp)
}

func (c *Config) overrideFromFlags() {

	// NOTE: override if flag value is not empty string
	c.PlatformURL = useOrDefault(PlatformURL(), c.PlatformURL)
	c.MessagingURL = useOrDefault(MessagingURL(), c.MessagingURL)
	c.RegistrationKey = useOrDefault(RegistrationKey(), c.RegistrationKey)
	c.SystemKey = useOrDefault(SystemKey(), c.SystemKey)
	c.SystemSecret = useOrDefault(SystemSecret(), c.SystemSecret)
	c.Developer.Email = useOrDefault(DeveloperEmail(), c.Developer.Email)
	c.Developer.Password = useOrDefault(DeveloperPassword(), c.Developer.Password)
	c.User.Email = useOrDefault(UserEmail(), c.User.Email)
	c.User.Password = useOrDefault(UserPassword(), c.User.Password)
	c.Device.Name = useOrDefault(DeviceName(), c.Device.Name)
	c.Device.ActiveKey = useOrDefault(DeviceActiveKey(), c.Device.ActiveKey)
	c.Out = useOrDefault(ConfigOut(), c.Out)

	// NOTE: let the boolean flag decide the value
	c.Import.ImportUsers = ShouldImportUsers()
	c.Import.ImportRows = ShouldImportRows()
}

// Config returns a *Config instance.
func (c *Config) Config(cbtest.T) *Config {
	return c
}

// HasSystem returns true if the given config has system information.
func (c *Config) HasSystem() bool {
	return len(c.SystemKey) > 0 && len(c.SystemSecret) > 0
}

// HasDeveloper returns true if the given config has developer information.
func (c *Config) HasDeveloper() bool {
	return c.Developer != nil && c.Developer.Email != "" && c.Developer.Password != ""
}

// HasUser returns true if the given config has user information.
func (c *Config) HasUser() bool {
	return c.User != nil && c.User.Email != "" && c.User.Password != ""
}

// HasDevice returns true if the given config has device information.
func (c *Config) HasDevice() bool {
	return c.Device != nil && c.Device.Name != "" && c.Device.ActiveKey != ""
}

// ShouldSave returns true if the out field of config is not empty (meaning the
// user to save the config to the specified path).
func (c *Config) ShouldSave() bool {
	return strings.TrimSpace(c.Out) != ""
}
