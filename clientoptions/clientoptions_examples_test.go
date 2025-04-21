package clientoptions_test

import (
	"context"
	"fmt"
	"net/url"

	"github.com/senzing-garage/go-grpcing/clientoptions"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleGetDialOptions_simple() {
	// For more information, visit
	// https://github.com/senzing-garage/go-grpcing/blob/main/clientoptions/clientoptions_test.go
	ctx := context.TODO()
	grpcURL := "grpc://localhost:8258"

	parsedURL, err := url.Parse(grpcURL)
	if err != nil {
		fmt.Println(err)
	}

	grpcDialOptions, err := clientoptions.GetDialOptions(ctx, *parsedURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(grpcDialOptions))
	// Output: 1
}
