[clearblade]: https://github.com/clearblade
[go-testing]: https://golang.org/pkg/testing/
[clearblade-go-sdk]: https://github.com/clearblade/Go-SDK

# cbtest

[![Go Report Card](https://goreportcard.com/badge/github.com/clearblade/cbtest)](https://goreportcard.com/report/github.com/clearblade/cbtest)

![License](https://img.shields.io/github/license/clearblade/cbtest)

*cbtest* is Go library that makes it easy to write automated tests against your
[ClearBlade][clearblade] systems. It provides the following features:

- Importing systems.

- Destroying systems.

- Combining systems before importing (base plus extra).

- Integrates with the Go testing package.

- Uses the [ClearBlade Go SDK][clearblade-go-sdk]

*cbtest* integrates with Golang's builtin [testing][go-testing] package rather
than trying to re-invent the wheel with a custom test suite. Check the
[examples/](examples/) folder to see what the tests look like.

## Similar tools

[terratest]: https://github.com/gruntwork-io/terratest

*cbtest* was inpired by [Terratest][terratest].

