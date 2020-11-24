package config

import (
	"flag"
	"strings"

	"github.com/clearblade/cbtest/internal/fsutil"
)

var (
	// Flags, registered during init.
	flagConfig          *string
	flagConfigOut       *string
	flagPlatformURL     *string
	flagMessagingURL    *string
	flagRegistrationKey *string
	flagSystemKey       *string
	flagSystemSecret    *string
	flagDevEmail        *string
	flagDevPassword     *string
	flagUserEmail       *string
	flagUserPassword    *string
	flagDeviceName      *string
	flagDeviceActiveKey *string
	flagImportUsers     *bool
	flagImportRows      *bool
)

// init registers our flags.
func init() {
	// NOTE: reminder that flag.Parse will be called by `go test`, so we don't need to call it here.
	flagConfig = flag.String("cbtest.config", "cbtest.json", "Path to the config file to use (credentials mostly)")
	flagConfigOut = flag.String("cbtest.config-out", "", "Path to write the config to")
	flagPlatformURL = flag.String("cbtest.platform-url", "", "Platform URL to use")
	flagMessagingURL = flag.String("cbtest.messaging-url", "", "Messaging URL to use")
	flagRegistrationKey = flag.String("cbtest.registration-key", "", "Registration key to use when creating developers")
	flagSystemKey = flag.String("cbtest.system-key", "", "System key to use")
	flagSystemSecret = flag.String("cbtest.system-secret", "", "System secret to use")
	flagDevEmail = flag.String("cbtest.dev-email", "", "Developer email to use")
	flagDevPassword = flag.String("cbtest.dev-password", "", "Developer password to use")
	flagUserEmail = flag.String("cbtest.user-email", "", "User email to use")
	flagUserPassword = flag.String("cbtest.user-password", "", "User password to use")
	flagDeviceName = flag.String("cbtest.device-name", "", "Device name to use")
	flagDeviceActiveKey = flag.String("cbtest.device-active-key", "", "Device active key to use")
	flagImportUsers = flag.Bool("cbtest.import-users", true, "Whenever users should be imported")
	flagImportRows = flag.Bool("cbtest.import-rows", true, "Whenever rows should be imported")
}

// FlagFound returns true if the flag with the given name was explicitly passed.
func FlagFound(name string) bool {

	if !flag.Parsed() {
		panic("FlagFound called before flag.Parse")
	}

	found := false

	flag.Visit(func(flag *flag.Flag) {
		if flag.Name == name {
			found = true
		}
	})

	return found
}

// ConfigPath returns the value given to the `-cbtest.config` flag.
func ConfigPath() string {

	if flagConfig == nil {
		panic("ConfigPath called before init")
	}

	if !flag.Parsed() {
		panic("ConfigPath called before flag.Parse")
	}

	return *flagConfig
}

// HasConfig returns true if a config path was provided.
func HasConfig() bool {
	return FlagFound("cbtest.config") || fsutil.IsFile(ConfigPath())
}

// ConfigOut returns the value given to the `-cbtest.config-out` flag.
func ConfigOut() string {
	if flagConfigOut == nil {
		panic("ConfigOut called before init")
	}

	if !flag.Parsed() {
		panic("ConfigOut called before flag.Parse")
	}

	return *flagConfigOut
}

// HasConfigOut returns true if the user passed the `-cbtest.config-out` flag.
func HasConfigOut() bool {
	configOut := ConfigOut()
	return strings.TrimSpace(configOut) != ""
}

// PlatformURL returns the value given to the `-cbtest.platform-url` flag.
func PlatformURL() string {

	if flagPlatformURL == nil {
		panic("PlatformURL called before init")
	}

	if !flag.Parsed() {
		panic("PlatformURL called before flag.Parse")
	}

	return *flagPlatformURL
}

// MessagingURL returns the value given to the `-cbtest.messaging-url` flag.
func MessagingURL() string {

	if flagMessagingURL == nil {
		panic("MessagingURL called before init")
	}

	if !flag.Parsed() {
		panic("MessagingURL called before flag.Parse")
	}

	return *flagMessagingURL
}

// RegistrationKey returns the value given to the `-cbtest.registration-key` flag.
func RegistrationKey() string {

	if flagRegistrationKey == nil {
		panic("RegistrationKey called before init")
	}

	if !flag.Parsed() {
		panic("RegistrationKey called before flag.Parse")
	}

	return *flagRegistrationKey
}

// SystemKey returns the value given to the `-cbtest.system-key` flag.
func SystemKey() string {

	if flagSystemKey == nil {
		panic("SystemKey called before init")
	}

	if !flag.Parsed() {
		panic("SystemKey called before flag.Parse")
	}

	return *flagSystemKey
}

// SystemSecret returns the value given to the `-cbtest.system-secret` flag.
func SystemSecret() string {

	if flagSystemSecret == nil {
		panic("SystemSecret called before init")
	}

	if !flag.Parsed() {
		panic("SystemSecret called before flag.Parse")
	}

	return *flagSystemSecret
}

// DeveloperEmail returns the value given to the `-cbtest.dev-email` flag.
func DeveloperEmail() string {

	if flagDevEmail == nil {
		panic("DeveloperEmail called before init")
	}

	if !flag.Parsed() {
		panic("DeveloperEmail called before flag.Parse")
	}

	return *flagDevEmail
}

// DeveloperPassword returns the value given to the `-cbtest.dev-password` flag.
func DeveloperPassword() string {

	if flagDevPassword == nil {
		panic("DeveloperPassword called before init")
	}

	if !flag.Parsed() {
		panic("DeveloperPassword called before flag.Parse")
	}

	return *flagDevPassword
}

// UserEmail returns the value given to the `-cbtest.user-email` flag.
func UserEmail() string {

	if flagUserEmail == nil {
		panic("UserEmail called before init")
	}

	if !flag.Parsed() {
		panic("UserEmail called before flag.Parse")
	}

	return *flagUserEmail
}

// UserPassword returns the value given to the `-cbtest.user-password` flag.
func UserPassword() string {

	if flagUserPassword == nil {
		panic("UserPassword called before init")
	}

	if !flag.Parsed() {
		panic("UserPassword called before flag.Parse")
	}

	return *flagUserPassword
}

// DeviceName returns the value given to the `-cbtest.device-name` flag.
func DeviceName() string {

	if flagDeviceName == nil {
		panic("DeviceName called before init")
	}

	if !flag.Parsed() {
		panic("DeviceName called before flag.Parse")
	}

	return *flagDeviceName
}

// DeviceActiveKey returns the value given to the `-cbtest.device-active-key` flag.
func DeviceActiveKey() string {

	if flagDeviceActiveKey == nil {
		panic("DeviceActiveKey called before init")
	}

	if !flag.Parsed() {
		panic("DeviceActiveKey called before flag.Parse")
	}

	return *flagDeviceActiveKey
}

// ShouldImportUsers returns the value given to the `-cbtest.import-users` flag.
func ShouldImportUsers() bool {

	if flagImportUsers == nil {
		panic("ShouldImportUsers called before init")
	}

	if !flag.Parsed() {
		panic("ShouldImportUsers called before flag.Parse")
	}

	return *flagImportUsers
}

// ShouldImportRows returns the value given to the `-cbtest.import-rows` flag.
func ShouldImportRows() bool {

	if flagImportRows == nil {
		panic("ShouldImportRows called before init")
	}

	if !flag.Parsed() {
		panic("ShouldImportRows called before flag.Parse")
	}

	return *flagImportRows
}
