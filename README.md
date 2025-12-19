# go-grpcing

If you are beginning your journey with [Senzing],
please start with [Senzing Quick Start guides].

You are in the [Senzing Garage] where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: template-go is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

Helpers for working with
[gRPC](https://grpc.io/).

[![Go Reference Badge]][Package reference]
[![Go Report Card Badge]][Go Report Card]
[![License Badge]][License]
[![go-test-linux.yaml Badge]][go-test-linux.yaml]
[![go-test-darwin.yaml Badge]][go-test-darwin.yaml]
[![go-test-windows.yaml Badge]][go-test-windows.yaml]

[![golangci-lint.yaml Badge]][golangci-lint.yaml]

## Overview

Examples of using `go-grpcing` can be seen in
[main.go](main.go)

## Use

```go
package main

import (
    "context"
    "fmt"
    "reflect"

    "github.com/senzing-garage/go-grpcing/grpcurl"
    "google.golang.org/grpc"
)

func main() {
    ctx := context.TODO()
    grpcUrl := "grpc://localhost:8258"

    grpcTarget, grpcDialOptions, err := grpcurl.Parse(ctx, grpcUrl)
    if err != nil {
        panic(err)
    }

    grpcConnection, err := grpc.Dial(grpcTarget, grpcDialOptions...)
    if err != nil {
        panic(err)
    }

    fmt.Printf("grpcTarget: %s\n", grpcTarget)
    fmt.Printf("grpcDialOptions: %s\n", reflect.TypeOf(grpcDialOptions))
    fmt.Printf("grpcConnection: %s\n", reflect.TypeOf(grpcConnection))
}
```

## References

1. [SDK documentation]
1. [Development]
1. [Errors]
1. [Examples]

[SDK documentation]: https://pkg.go.dev/github.com/senzing-garage/go-grpcing
[Development]: docs/development.md
[Errors]: docs/errors.md
[Examples]: docs/examples.md
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/template-go.svg
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/senzing-garage/template-go
[Go Report Card]: https://goreportcard.com/report/github.com/senzing-garage/template-go
[go-test-darwin.yaml Badge]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-darwin.yaml/badge.svg
[go-test-darwin.yaml]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-darwin.yaml
[go-test-linux.yaml Badge]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-linux.yaml/badge.svg
[go-test-linux.yaml]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-linux.yaml
[go-test-windows.yaml Badge]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-windows.yaml/badge.svg
[go-test-windows.yaml]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-windows.yaml
[golangci-lint.yaml Badge]: https://github.com/senzing-garage/template-go/actions/workflows/golangci-lint.yaml/badge.svg
[golangci-lint.yaml]: https://github.com/senzing-garage/template-go/actions/workflows/golangci-lint.yaml
[License Badge]: https://img.shields.io/badge/License-Apache2-brightgreen.svg
[License]: https://github.com/senzing-garage/template-go/blob/main/LICENSE
[Package reference]: https://pkg.go.dev/github.com/senzing-garage/template-go
[Senzing Garage]: https://github.com/senzing-garage
[Senzing Quick Start guides]: https://docs.senzing.com/quickstart/
[Senzing]: https://senzing.com/
