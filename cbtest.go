package cbtest

import (
	"flag"
)

var (
	// Flags, registered during init.
	flagConfig *string
)

// init registers our flags.
func init() {
	// NOTE: reminder that flag.Parse will be called by `go test`, so we don't need to call it here.
	flagConfig = flag.String("cbtest.config", "cbtest.yml", "Path to the config file to use (credentials mostly)")
}

// ConfigPath the value given to the `-cbtest.config` flag.
func ConfigPath() string {

	if flagConfig == nil {
		panic("ConfigPath called before init")
	}

	if !flag.Parsed() {
		panic("ConfigPath called flag.Parse")
	}

	return *flagConfig
}
