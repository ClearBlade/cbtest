// Package modules contains all the cbtest core modules.
//
// Most core modules expose functions that will either fail a test or return
// an error:
//
//     module.DoSomething(t, ...) // will panic on failure
//     ...
//     module.DoSomethingE(t, ...) // will return error on failure.
//
package modules
