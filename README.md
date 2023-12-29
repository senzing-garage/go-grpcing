# go-grpcing

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## Synopsis

Helpers for working with
[gRPC](https://grpc.io/).

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/go-grpcing.svg)](https://pkg.go.dev/github.com/senzing-garage/go-grpcing)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/go-grpcing)](https://goreportcard.com/report/github.com/senzing-garage/go-grpcing)
[![go-test.yaml](https://github.com/senzing-garage/go-grpcing/actions/workflows/go-test.yaml/badge.svg)](https://github.com/senzing-garage/go-grpcing/actions/workflows/go-test.yaml)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/go-grpcing/blob/main/LICENSE)

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

1. [API documentation](https://pkg.go.dev/github.com/senzing-garage/go-grpcing)
1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
