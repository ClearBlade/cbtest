package cbtest

import (
	"encoding/json"
	"io"
	"os"
)

// Config contains information about the system that is gonna be used for
// testing.
type Config struct {
	PlatformURL     string     `json:"platformUrl" mapstructure:"platformUrl"`
	MessagingURL    string     `json:"messagingUrl" mapstructure:"messagingUrl"`
	RegistrationKey string     `json:"registrationKey,omitempty" mapstructure:"registrationKey"`
	SystemKey       string     `json:"systemKey,omitempty" mapstructure:"systemKey"`
	SystemSecret    string     `json:"systemSecret,omitempty" mapstructure:"systemSecret"`
	Developer       *Developer `json:"developer,omitempty" mapstructure:"developer"`
}

// Developer contains the developer credentials that must be provided if using
// an existing system.
type Developer struct {
	Email    string `json:"email" mapstructure:"email"`
	Password string `json:"password" mapstructure:"password"`
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

	config := Config{}
	err := json.NewDecoder(r).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}

// HasSystem returns true if the given config has system information.
func (c *Config) HasSystem() bool {
	return len(c.SystemKey) > 0 && len(c.SystemSecret) > 0
}

// HasDeveloper returns true if the given config has developer information.
func (c *Config) HasDeveloper() bool {
	return c.Developer != nil && c.Developer.Email != "" && c.Developer.Password != ""
}
