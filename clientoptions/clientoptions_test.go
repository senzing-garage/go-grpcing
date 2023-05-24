package clientoptions

import (
	"context"
	"errors"
	"fmt"
	"net/url"
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
	expectedError          error
}{
	{
		name:                   "clientoptions-0001",
		url:                    "grpc://localhost",
		expectedAddress:        "localhost",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
		expectedError:          nil,
	},
	{
		name:                   "clientoptions-0002",
		url:                    "grpc://localhost:1234",
		expectedAddress:        "localhost:1234",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
		expectedError:          nil,
	},
	{
		name:                   "clientoptions-0003",
		url:                    `http://localhost:1234`,
		expectedAddress:        "",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
		expectedError:          nil,
	},
	{
		name:                   "clientoptions-0004",
		url:                    `grpc://localhost:1234/bob/?something="bob2"`,
		expectedAddress:        "localhost:1234",
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

func TestGetDialOptions(test *testing.T) {
	ctx := context.TODO()
	for _, testCase := range testCasesForGrpcurl {
		test.Run(testCase.name, func(test *testing.T) {
			parsedUrl, err := url.Parse(testCase.url)
			assert.Nil(test, err)
			grpcOptions, err := GetDialOptions(ctx, *parsedUrl)
			assert.Equal(test, testCase.expectedError, err, testCase.name+"-err")
			assert.Equal(test, testCase.expectedDialOptionsLen, len(grpcOptions), testCase.name+"-GrpcOptionsLen")
			// assert.Equal(test, testCase.expectedDialOptions, grpcOptions, testCase.name+"-GrpcOptions")
		})
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleGetDialOptions_simple() {
	// For more information, visit https://github.com/Senzing/go-grpcing/blob/main/clientoptions/clientoptions_test.go
	ctx := context.TODO()
	grpcUrl := "grpc://localhost:8258"
	parsedUrl, err := url.Parse(grpcUrl)
	if err != nil {
		fmt.Println(err)
	}
	grpcOptions, err := GetDialOptions(ctx, *parsedUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(grpcOptions))
	// Output: 1
}
