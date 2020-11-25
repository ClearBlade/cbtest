// Package lookup provides functions for (unsafely) looking up values in dynamic
// maps, structures, etc.
//
// Module is powered by:
// https://github.com/mcuadros/go-lookup
//
// Note that this package is meant for testing purposes only. On a real application
// you would avoid unsafe reading of maps and structs as much as possible, and try
// to use the type system to your advantage.
package lookup

import (
	"github.com/mcuadros/go-lookup"
	"github.com/stretchr/testify/require"

	"github.com/clearblade/cbtest"
)

// Path performs a lookup into the given value using a string path.
// Panics on failure.
func Path(t cbtest.T, i interface{}, path string) interface{} {
	res, err := PathE(t, i, path)
	require.NoError(t, err)
	return res
}

// PathE performs a lookup into the given value using a string path.
// Returns error on failure.
func PathE(t cbtest.T, i interface{}, path string) (interface{}, error) {
	value, err := lookup.LookupString(i, path)
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// PathI performs a lookup into the given value using a case-insensitive string path.
// Panics on failure.
func PathI(t cbtest.T, i interface{}, path string) interface{} {
	res, err := PathIE(t, i, path)
	require.NoError(t, err)
	return res
}

// PathIE performs a lookup into the given value using a case-insensitive string path.
// Returns error on failure.
func PathIE(t cbtest.T, i interface{}, path string) (interface{}, error) {
	value, err := lookup.LookupStringI(i, path)
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}
