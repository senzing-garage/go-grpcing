package grpcurl

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var testCasesForGrpcurl = []struct {
	name                   string
	url                    string
	expectedAddress        string
	expectedDialOptions    []grpc.DialOption
	expectedDialOptionsLen int
}{
	{
		name:                   "parser-0001",
		url:                    "grpc://localhost:1234",
		expectedAddress:        "localhost:1234",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
	},
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestParse(test *testing.T) {
	ctx := context.TODO()
	for _, testCase := range testCasesForGrpcurl {
		test.Run(testCase.name, func(test *testing.T) {
			grpcAddress, grpcOptions, err := Parse(ctx, testCase.url)
			assert.Nil(test, err, testCase.name+"-err")
			assert.Equal(test, testCase.expectedAddress, grpcAddress, testCase.name+"-GrpcAddress")
			assert.Equal(test, testCase.expectedDialOptionsLen, len(grpcOptions), testCase.name+"-GrpcOptionsLen")
			// assert.Equal(test, testCase.expectedDialOptions, grpcOptions, testCase.name+"-GrpcOptions")

		})
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleParse_simple() {
	// For more information, visit https://github.com/Senzing/go-grpcing/blob/main/examplepackage/examplepackage_test.go
	ctx := context.TODO()
	grpcUrl := "grpc://localhost:8258"
	grpcAddress, grpcOptions, err := Parse(ctx, grpcUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(grpcAddress, len(grpcOptions))
	// Output:
	// localhost:8258 1
}
