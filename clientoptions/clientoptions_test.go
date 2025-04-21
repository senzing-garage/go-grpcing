package clientoptions_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/senzing-garage/go-grpcing/clientoptions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		name:                   "clientoptions-0001",
		url:                    "grpc://localhost",
		expectedTarget:         "localhost",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
		expectedError:          nil,
	},
	{
		name:                   "clientoptions-0002",
		url:                    "grpc://localhost:1234",
		expectedTarget:         "localhost:1234",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
		expectedError:          nil,
	},
	{
		name:                   "clientoptions-0003",
		url:                    `http://localhost:1234`,
		expectedTarget:         "",
		expectedDialOptions:    []grpc.DialOption{},
		expectedDialOptionsLen: 1,
		expectedError:          nil,
	},
	{
		name:                   "clientoptions-0004",
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

func TestGetDialOptions(test *testing.T) {
	test.Parallel()
	ctx := test.Context()

	for _, testCase := range testCasesForGrpcurl {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			parsedURL, err := url.Parse(testCase.url)
			require.NoError(test, err)
			grpcDialOptions, err := clientoptions.GetDialOptions(ctx, *parsedURL)
			assert.Equal(test, testCase.expectedError, err, testCase.name+"-err")
			assert.Len(test, grpcDialOptions, testCase.expectedDialOptionsLen, testCase.name+"-DialOptionsLen")
		})
	}
}
