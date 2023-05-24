/*
 */
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

	grpcAddress, grpcOptions, err := grpcurl.Parse(ctx, grpcUrl)
	if err != nil {
		panic(err)
	}

	grpcConnection, err := grpc.Dial(grpcAddress, grpcOptions...)
	if err != nil {
		panic(err)
	}

	fmt.Printf("grpcAddress: %s\n", grpcAddress)
	fmt.Printf("grpcOptions: %s\n", reflect.TypeOf(grpcOptions))
	fmt.Printf("grpcConnection: %s\n", reflect.TypeOf(grpcConnection))
}
