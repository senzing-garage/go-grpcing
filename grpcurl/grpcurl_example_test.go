package grpcurl_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-grpcing/grpcurl"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleParse_simple() {
	// For more information, visit https://github.com/senzing-garage/go-grpcing/blob/main/grpcurl/grpcurl_test.go
	ctx := context.TODO()
	grpcURL := "grpc://localhost:8258"

	grpcTarget, grpcDialOptions, err := grpcurl.Parse(ctx, grpcURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(grpcTarget, len(grpcDialOptions))
	// Output:
	// localhost:8258 1
}
