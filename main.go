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
