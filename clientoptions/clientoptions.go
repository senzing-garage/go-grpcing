package clientoptions

import (
	"context"
	"fmt"
	"net/url"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
The GetDialOptions function returns a slice of grpc.DialOption(s) from a parsed URL.

Input
  - ctx: A context to control lifecycle.
  - parsedUrl: A parsed URL.

Output
  - grpcDialOptions: A slice of configuration options for gRPC server/client
*/
func GetDialOptions(ctx context.Context, parsedURL url.URL) ([]grpc.DialOption, error) {
	_ = ctx
	var err error
	var grpcDialOptions = []grpc.DialOption{}

	// userName := parsedUrl.User.Username()
	// password, isSet := parsedUrl.User.Password()

	queryParameters := parsedURL.Query()
	if len(queryParameters) == 0 {
		grpcDialOptions = append(grpcDialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		err = fmt.Errorf("not sure how to parse gRPC URL: %s", parsedURL.String())
	}

	return grpcDialOptions, err
}
