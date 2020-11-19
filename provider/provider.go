// Package provider defines interface for obtaining instances of objects that
// we commonly use in all the modules. This pattern is useful to make our
// modules more flexible, as they don't take concrete objects but interfaces,
// which can be mocked, replaced, etc.
package provider

import (
	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest"
	c "github.com/clearblade/cbtest/config"
)

// Config provides a *config.Config instance.
type Config interface {
	Config(cbtest.T) *c.Config
}

// Client provides a *cb.DevClient instance. Note that this is an operation that
// can fail, therefore, we have a function that panics and another one that returns
// the error.
type Client interface {
	Client(cbtest.T) *cb.DevClient
	ClientE(cbtest.T) (*cb.DevClient, error)
}

// ConfigAndClient provides both provider.Config and provider.Client.
type ConfigAndClient interface {
	Config
	Client
}
