[danielrs]: https://github.com/danielrs
[clearblade]: https://github.com/clearblade
[go-testing]: https://golang.org/pkg/testing/
[clearblade-go-sdk]: https://github.com/clearblade/Go-SDK

**Disclaimer**: this project is currently a pet project by [danielrs][danielrs].

# cbtest

[![Go Report Card](https://goreportcard.com/badge/github.com/clearblade/cbtest)](https://goreportcard.com/report/github.com/clearblade/cbtest)
![GitHub last commit](https://img.shields.io/github/last-commit/clearblade/cbtest)
![GitHub license](https://img.shields.io/github/license/clearblade/cbtest)

`cbtest` is Go library that makes it easy to write automated tests against your
[ClearBlade][clearblade] systems. It provides the following features:

- Integrates with the [Go testing][go-testing] package.

- Importing systems.

- Destroying systems.

- Combining systems before importing (base system plus extra).

- Uses the [ClearBlade Go SDK][clearblade-go-sdk].

## Examples

Check the [examples/](examples/) folder for reference.

## Similar tools

[terratest]: https://github.com/gruntwork-io/terratest

`cbtest` was inpired by [Terratest][terratest].
