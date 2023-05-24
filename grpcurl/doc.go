/*
grpcurl is used to transfer information conveyed in a URL into an address of a gRPC server and a slice of grpc.DialOption.

# Overview

More information at https://github.com/senzing/go-grpcing

# Examples

Example of use:

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
	}

Additional examples of use can be seen in https://github.com/Senzing/go-grpcing/blob/main/grpcurl/grpcurl_test.go
*/
package grpcurl
