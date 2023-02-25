# CloudZero Client Go

[![Testing](https://github.com/jjttech/cloudzero-client-go/workflows/Testing/badge.svg)](https://github.com/jjttech/cloudzero-client-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jjttech/cloudzero-client-go?style=flat-square)](https://goreportcard.com/report/github.com/jjttech/cloudzero-client-go)
[![GoDoc](https://godoc.org/github.com/jjttech/cloudzero-client-go?status.svg)](https://godoc.org/github.com/jjttech/cloudzero-client-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/jjttech/cloudzero-client-go/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/jjttech/cloudzero-client-go/all.svg)](https://github.com/jjttech/cloudzero-client-go/releases/latest)

## Example

See [examples](examples/):

```go
package main

import (
  "github.com/jjttech/cloudzero-client-go/cloudzero"
  log "github.com/sirupsen/logrus"
)

func main() {
  var (
    cz  *cloudzero.CloudZero
    err error
  )

  if cz, err = cloudzero.New(); err != nil {
    log.WithError(err).Fatal("unable to create CloudZero client")
  }

  // Load from the default filename "definition.yaml" in the current directory
  def, err := cz.CostFormation.Read(cloudzero.DefaultDefinitionFilename)
  if err != nil {
    log.WithError(err).Fatal("unable to load file")
  }

  // Print to the screen
  if err = cz.CostFormation.Write(def, ""); err != nil {
    log.WithError(err).Fatal("unable to write file")
  }
}
```

## Development

### Requirements

* Go 1.19.0+
* GNU Make
* git


### Building

```
# Default target is 'build'
$ make

# Explicitly run build
$ make build

# Locally test the CI build scripts
# make build-ci
```


### Testing

Before contributing, all linting and tests must pass.  Tests can be run directly via:

```
# Tests and Linting
$ make test

# Only unit tests
$ make test-unit

# Only integration tests
$ make test-integration
```

### Commit Messages

Using the following format for commit messages allows for auto-generation of
the [CHANGELOG](CHANGELOG.md):

#### Format:

`<type>(<scope>): <subject>`

| Type | Description | Change log? |
|------| ----------- | :---------: |
| `chore` | Maintenance type work | No |
| `docs` | Documentation Updates | Yes |
| `feat` | New Features | Yes |
| `fix`  | Bug Fixes | Yes |
| `refactor` | Code Refactoring | No |

#### Scope

This refers to what part of the code is the focus of the work.  For example:

**General:**

* `build` - Work related to the build system (linting, makefiles, CI/CD, etc)
* `release` - Work related to cutting a new release

**Package Specific:**

* `cmd/jjttech` - Work related to the jjttech binary / command
* `internal/http` - Work related to the `internal/http` package
* `pkg/alerts` - Work related to the `pkg/alerts` package



### Documentation

**Note:** This requires the repo to be in your GOPATH [(godoc issue)](https://github.com/golang/go/issues/26827)

```
$ make docs
```

## Open Source License

This project is distributed under the [Apache 2 license](LICENSE).
