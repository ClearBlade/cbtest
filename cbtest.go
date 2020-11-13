package cbtest

import (
	"flag"
)

var (
	// Flags, registered during init.
	flagConfig       *string
	flagPlatformURL  *string
	flagMessagingURL *string
	flagSystemKey    *string
	flagSystemSecret *string
	flagDevEmail     *string
	flagDevPassword  *string
	flagUserEmail    *string
	flagUserPassword *string
	flagImportUsers  *bool
	flagImportRows   *bool
	flagKeep         *bool
)

// init registers our flags.
func init() {
	// NOTE: reminder that flag.Parse will be called by `go test`, so we don't need to call it here.
	flagConfig = flag.String("cbtest.config", "cbtest.yml", "Path to the config file to use (credentials mostly)")
	flagPlatformURL = flag.String("cbtest.platform-url", "", "Platform URL to use")
	flagMessagingURL = flag.String("cbtest.messaging-url", "", "Messaging URL to use")
	flagSystemKey = flag.String("cbtest.system-key", "", "System key to use")
	flagSystemSecret = flag.String("cbtest.system-secret", "", "System secret to use")
	flagDevEmail = flag.String("cbtest.dev-email", "", "Developer email to use")
	flagDevPassword = flag.String("cbtest.dev-password", "", "Developer password to use")
	flagUserEmail = flag.String("cbtest.user-email", "", "User email to use")
	flagUserPassword = flag.String("cbtest.user-password", "", "User password to use")
	flagImportUsers = flag.Bool("cbtest.import-users", true, "Whenever users should be imported")
	flagImportRows = flag.Bool("cbtest.import-rows", true, "Whenever rows should be imported")
	flagKeep = flag.Bool("cbtest.keep", false, "Keep systems that were created by cbtest (external systems are never destroyed)")
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

// ShouldImportRows returns the value given to the `-cbtest.import-users` flag.
func ShouldImportRows() bool {

	if flagImportRows == nil {
		panic("ShouldImportRows called before init")
	}

	if !flag.Parsed() {
		panic("ShouldImportRows called before flag.Parse")
	}

	return *flagImportRows
}

// ShouldKeepSystem returns the value given to the `-cbtest.keep` flag.
func ShouldKeepSystem() bool {

	if flagKeep == nil {
		panic("ShouldKeepSystem called before init")
	}

	if !flag.Parsed() {
		panic("ShouldKeepSystem called before flag.Parse")
	}

	return *flagKeep
}
