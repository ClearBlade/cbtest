package cbtest

import (
	"flag"
)

var (
	// Flags, registered during init.
	flagConfig       *string
	flagPlatformURL  *string
	flagMessagingURL *string
	flagDevEmail     *string
	flagDevPassword  *string
	flagUserEmail    *string
	flagUserPassword *string
	flagKeep         *bool
)

// init registers our flags.
func init() {
	// NOTE: reminder that flag.Parse will be called by `go test`, so we don't need to call it here.
	flagConfig = flag.String("cbtest.config", "cbtest.yml", "Path to the config file to use (credentials mostly)")
	flagPlatformURL = flag.String("cbtest.platform-url", "", "Platform URL to use")
	flagMessagingURL = flag.String("cbtest.messaging-url", "", "Messaging URL to use")
	flagDevEmail = flag.String("cbtest.dev-email", "", "Developer email to use")
	flagDevPassword = flag.String("cbtest.dev-password", "", "Developer password to use")
	flagUserEmail = flag.String("cbtest.user-email", "", "User email to use")
	flagUserPassword = flag.String("cbtest.user-password", "", "User password to use")
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
