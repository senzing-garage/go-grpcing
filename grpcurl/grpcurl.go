package grpcurl

import (
	"context"
	"fmt"
	"net/url"

	"github.com/senzing-garage/go-grpcing/clientoptions"
	"google.golang.org/grpc"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
The Parse function parses a URL to extract information to create GrpcTarget and
GrpcDialOptions.

Input
  - ctx: A context to control lifecycle.
  - grpcUrl: A URL starting with grpc://

Output
  - grpcTarget: The hostname:port
  - grpcDialOptions: A slice of configuration options for gRPC server/client
*/
func Parse(ctx context.Context, grpcURL string) (string, []grpc.DialOption, error) {
	var grpcTarget = ""
	var grpcDialOptions = []grpc.DialOption{}

	parsedURL, err := url.Parse(grpcURL)
	if err != nil {
		return grpcTarget, grpcDialOptions, err
	}

	switch parsedURL.Scheme {
	case "grpc":
		grpcTarget = parsedURL.Host
		grpcDialOptions, err = clientoptions.GetDialOptions(ctx, *parsedURL)
	default:
		err = fmt.Errorf("gRPC URL must start with grpc://, not %s://.  (%s)", parsedURL.Scheme, grpcURL)
	}
	return grpcTarget, grpcDialOptions, err
}
