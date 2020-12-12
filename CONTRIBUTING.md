[go-testing]: https://golang.org/pkg/testing/

First of all, thank you for being interested in contributing to `cbtest`.

### TL; DR

1. **Every** change not comming from the core developers should go through a PR.

2. **Every** change or bug fix should also include a unit test.

3. New modules go into the `contrib` directory.

4. Core modules in the `modules` directory need to be thoroughly tested.

5. All modules follow the same conventions (see below).

Code that breaks the guidelines above will be rejected.

----

### Overview

`cbtest` is being developed with the following points in mind:

1. Easy to get started.

2. Provide modules that are easier to locate and navigate (`contrib` or
   `modules` packages).

3. Provide helper functions and assertions that make tests more expressive.

4. Make use of the Golang [testing][go-testing] library whenever possible.

5. New features must include tests.

### Changing the config format, flags, etc

The `config` package defines the types and functions for reading the config, as
well as the `cbtest.*` flags that we can pass. As a convention, config values
will only be read from the outside world from the `config` package, all the
other modules will treat the config as immutable once created.

When adding new flags to `cbtest` all of them should start with the prefix
`flag*` (because flags are package global).

### Adding new modules

Most new modules will start in the `contrib` package, once they are in good shape,
they will be moved to the `modules` package. Modules will usually expose functions
that:

1. Make the interaction with the ClearBade system easier (query data, etc).

2. Assert the state of the ClearBlade system or interaction (collection, service response, etc).

#### Function signatures

Each exposed function in the modules must provide the follow variants:

```Go
...
[Fname](t cbtest.T, provider provider.ConfigAndClient[, rest args]) Rtype // panic on failure
...
[Fname]E(t cbtest.T, provider provider.ConfigAndClient[, rest args]) (Rtype, error) // return error on failure
...
```

1. Takes the `cbtest.T` interface as first argument.

2. Takes a `provider.ConfigAndClient` as second argument.

3. Takes zero or more arguments specific to the function being implemented.

4. If the function is the non-E variant, it panics on failure.

5. If the function is the E variant, it returns error on failure.
