# go-grpcing

## Synopsis

Helpers for working with
[gRPC](https://grpc.io/).

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-grpcing.svg)](https://pkg.go.dev/github.com/senzing/go-grpcing)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/go-grpcing)](https://goreportcard.com/report/github.com/senzing/go-grpcing)
[![go-test.yaml](https://github.com/Senzing/go-grpcing/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/go-grpcing/actions/workflows/go-test.yaml)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/Senzing/go-grpcing/blob/main/LICENSE)

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

    "github.com/senzing/go-grpcing/grpcurl"
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

1. [API documentation](https://pkg.go.dev/github.com/senzing/go-grpcing)
1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
