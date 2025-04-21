/*
 */
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
	grpcURL := "grpc://localhost:8258"

	grpcTarget, grpcDialOptions, err := grpcurl.Parse(ctx, grpcURL)
	if err != nil {
		panic(err)
	}

	grpcConnection, err := grpc.NewClient(grpcTarget, grpcDialOptions...)
	if err != nil {
		panic(err)
	}

	outputf("grpcTarget: %s\n", grpcTarget)
	outputf("grpcDialOptions: %s\n", reflect.TypeOf(grpcDialOptions))
	outputf("grpcConnection: %s\n", reflect.TypeOf(grpcConnection))
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func outputf(format string, message ...any) {
	fmt.Printf(format, message...) //nolint
}
