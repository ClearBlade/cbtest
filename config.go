package cbtest

import (
	"encoding/json"
	"io"
	"os"
)

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

	config.overrideFromFlags()
	return config, err
}

func (c *Config) overrideFromFlags() {
	c.PlatformURL = useOrDefault(PlatformURL(), c.PlatformURL)
	c.MessagingURL = useOrDefault(MessagingURL(), c.MessagingURL)
	c.SystemKey = useOrDefault(SystemKey(), c.SystemKey)
	c.SystemSecret = useOrDefault(SystemSecret(), c.SystemSecret)
	c.Developer.Email = useOrDefault(DeveloperEmail(), c.Developer.Email)
	c.Developer.Password = useOrDefault(DeveloperPassword(), c.Developer.Password)
	c.User.Email = useOrDefault(UserEmail(), c.User.Email)
	c.User.Password = useOrDefault(UserPassword(), c.User.Password)
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
