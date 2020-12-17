[danielrs]: https://github.com/danielrs
[clearblade]: https://github.com/clearblade
[go-testing]: https://golang.org/pkg/testing/
[clearblade-go-sdk]: https://github.com/clearblade/Go-SDK

**Disclaimer**: `cbtest` is currently a pet project by [danielrs][danielrs].

# cbtest

[![PkgGoDev](https://pkg.go.dev/badge/github.com/clearblade/cbtest)](https://pkg.go.dev/github.com/clearblade/cbtest)
[![Go Report Card](https://goreportcard.com/badge/github.com/clearblade/cbtest)](https://goreportcard.com/report/github.com/clearblade/cbtest)
[![Test status](https://github.com/clearblade/cbtest/workflows/tests/badge.svg?branch=master "test status")](https://github.com/clearblade/cbtest/actions)
![GitHub last commit](https://img.shields.io/github/last-commit/clearblade/cbtest)
[![GitHub license](https://img.shields.io/github/license/clearblade/cbtest)](LICENSE)

`cbtest` is Go library that makes it easy to write automated tests against your
[ClearBlade][clearblade] systems. It provides the following features:

- Integrates with the [Go testing][go-testing] package and similar libraries.

- Importing systems (with multiple system merge).

- Destroying systems.

- Using existing systems.

- Uses the [ClearBlade Go SDK][clearblade-go-sdk].

## Examples

They are located under the [examples](examples/) folder. You can run any of
them like follows:

```bash
go test -v ./examples/[EXAMPLE FOLDER]
```

## Wiki

[wiki]: https://github.com/ClearBlade/cbtest/wiki

Please refer to our [wiki][wiki] for more details regarding `cbtest`.

## Contributing

To contribute, please follow the [guidelines](CONTRIBUTING.md).

## References

[advanced-testing-in-go]: https://about.sourcegraph.com/go/advanced-testing-in-go/
[terratest]: https://github.com/gruntwork-io/terratest

- `cbtest` was inpired by [Terratest][terratest].
- `cbtest` uses many ideas from [Advanced Testing in Go][advanced-testing-in-go].
