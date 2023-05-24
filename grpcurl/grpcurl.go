package grpcurl

import (
	"context"
	"fmt"
	"net/url"

	"github.com/senzing/go-grpcing/clientoptions"
	"google.golang.org/grpc"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
The Parse function parses a URL to extract information to create GrpcAddress and
GrpcOptions.

Input
  - ctx: A context to control lifecycle.
  - grpcUrl: A URL starting with grpc://

Output
  - grpcAddress: The hostname:port
  - grpcDialOptions: A list of configuration options for gRPC server/client
*/
func Parse(ctx context.Context, grpcUrl string) (string, []grpc.DialOption, error) {
	var err error = nil
	var grpcAddress = ""
	var grpcOptions = []grpc.DialOption{}

	parsedUrl, err := url.Parse(grpcUrl)
	if err != nil {
		return grpcAddress, grpcOptions, err
	}

	switch parsedUrl.Scheme {
	case "grpc":
		grpcAddress = parsedUrl.Host
		grpcOptions, err = clientoptions.GetDialOptions(ctx, *parsedUrl)
	default:
		err = fmt.Errorf("gRPC URL must start with grpc://, not %s://.  (%s)", parsedUrl.Scheme, grpcUrl)
	}
	return grpcAddress, grpcOptions, err
}
