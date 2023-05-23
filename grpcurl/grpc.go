package grpcurl

import (
	"context"
	"fmt"
	"net/url"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func parseGrpc(ctx context.Context, parsedUrl url.URL) (string, []grpc.DialOption, error) {
	var err error = nil
	var grpcAddress = ""
	var grpcOptions = []grpc.DialOption{}

	grpcAddress = fmt.Sprintf("%s:%s", parsedUrl.Hostname(), parsedUrl.Port())

	// userName := parsedUrl.User.Username()
	// password, isSet := parsedUrl.User.Password()

	queryParameters := parsedUrl.Query()
	if len(queryParameters) == 0 {
		grpcOptions = append(grpcOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		err = fmt.Errorf("not sure how to parse gRPC URL: %s", parsedUrl.String())
	}

	return grpcAddress, grpcOptions, err
}

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
		grpcAddress, grpcOptions, err = parseGrpc(ctx, *parsedUrl)
	default:
		err = fmt.Errorf("gRPC URL must start with grpc://, not %s://.  (%s)", parsedUrl.Scheme, grpcUrl)
	}
	return grpcAddress, grpcOptions, err
}
