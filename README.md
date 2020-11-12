[clearblade]: https://github.com/clearblade
[go-testing]: https://golang.org/pkg/testing/

# cbtest

*cbtest* is Go library that makes it easy to write automated tests against your
[ClearBlade][clearblade] systems. It provides the following features:

- Importing systems.

- Destroying systems.

- Combining systems before importing (base plus extra).

*cbtest* integrates with Golang's builtin [testing][go-testing] package rather
than trying to re-invent the wheel with a custom test suite. Check the
[examples/](examples/) folder to see what the tests look like.

## Similar tools

[terratest]: https://github.com/gruntwork-io/terratest

*cbtest* was inpired by [Terratest][terratest].

