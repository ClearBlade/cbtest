package cbtest

import (
	"flag"
)

var (
	// Flags, registered during init.
	flagConfig *string
	flagKeep   *bool
)

// init registers our flags.
func init() {
	// NOTE: reminder that flag.Parse will be called by `go test`, so we don't need to call it here.
	flagConfig = flag.String("cbtest.config", "cbtest.yml", "Path to the config file to use (credentials mostly)")
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
