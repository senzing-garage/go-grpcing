package grpcurl_test

import (
	"fmt"
	"testing"

	"github.com/senzing-garage/go-grpcing/grpcurl"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var err error

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
		expectedError: fmt.Errorf(
			"gRPC URL must start with grpc://, not http://.  (http://localhost:1234) Error: %w",
			err,
		),
	},
	{
		name:                   "grpcurl-0004",
		url:                    `grpc://localhost:1234/bob/?something="bob2"`,
		expectedTarget:         "localhost:1234",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 0,
		expectedError: fmt.Errorf(
			"not sure how to parse gRPC URL: grpc://localhost:1234/bob/?something=\"bob2\" Error: %w", err,
		),
	},
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestParse(test *testing.T) {
	test.Parallel()
	ctx := test.Context()

	for _, testCase := range testCasesForGrpcurl {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			grpcTarget, grpcDialOptions, err := grpcurl.Parse(ctx, testCase.url)
			assert.Equal(test, testCase.expectedError, err, testCase.name+"-err")
			assert.Equal(test, testCase.expectedTarget, grpcTarget, testCase.name+"-GrpcTarget")
			assert.Len(test, grpcDialOptions, testCase.expectedDialOptionsLen, testCase.name+"-GrpcDialOptionsLen")
		})
	}
}
