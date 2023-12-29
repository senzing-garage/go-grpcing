package grpcurl

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var testCasesForGrpcurl = []struct {
	name                   string
	url                    string
	expectedTarget         string
	expectedDialOptions    []grpc.DialOption
	expectedDialOptionsLen int
	expectedError          error
}{
	{
		name:                   "grpcurl-0001",
		url:                    "grpc://localhost",
		expectedTarget:         "localhost",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
		expectedError:          nil,
	},
	{
		name:                   "grpcurl-0002",
		url:                    "grpc://localhost:1234",
		expectedTarget:         "localhost:1234",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
		expectedError:          nil,
	},
	{
		name:                   "grpcurl-0003",
		url:                    `http://localhost:1234`,
		expectedTarget:         "",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 0,
		expectedError:          errors.New("gRPC URL must start with grpc://, not http://.  (http://localhost:1234)"),
	},
	{
		name:                   "grpcurl-0004",
		url:                    `grpc://localhost:1234/bob/?something="bob2"`,
		expectedTarget:         "localhost:1234",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 0,
		expectedError:          errors.New("not sure how to parse gRPC URL: grpc://localhost:1234/bob/?something=\"bob2\""),
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
			grpcTarget, grpcDialOptions, err := Parse(ctx, testCase.url)
			assert.Equal(test, testCase.expectedError, err, testCase.name+"-err")
			assert.Equal(test, testCase.expectedTarget, grpcTarget, testCase.name+"-GrpcTarget")
			assert.Equal(test, testCase.expectedDialOptionsLen, len(grpcDialOptions), testCase.name+"-GrpcDialOptionsLen")
			// assert.Equal(test, testCase.expectedDialOptions, grpcDialOptions, testCase.name+"-GrpcDialOptions")
		})
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleParse_simple() {
	// For more information, visit https://github.com/senzing-garage/go-grpcing/blob/main/grpcurl/grpcurl_test.go
	ctx := context.TODO()
	grpcUrl := "grpc://localhost:8258"
	grpcTarget, grpcDialOptions, err := Parse(ctx, grpcUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(grpcTarget, len(grpcDialOptions))
	// Output:
	// localhost:8258 1
}
