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
The GetDialOptions function generates a slice of grpc.DialOption(s) from a parsed URL.

Input
  - ctx: A context to control lifecycle.
  - parsedUrl: A parsed URL.

Output
  - grpcDialOptions: A list of configuration options for gRPC server/client
*/
func GetDialOptions(ctx context.Context, parsedUrl url.URL) ([]grpc.DialOption, error) {
	var err error = nil
	var grpcOptions = []grpc.DialOption{}

	// userName := parsedUrl.User.Username()
	// password, isSet := parsedUrl.User.Password()

	queryParameters := parsedUrl.Query()
	if len(queryParameters) == 0 {
		grpcOptions = append(grpcOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		err = fmt.Errorf("not sure how to parse gRPC URL: %s", parsedUrl.String())
	}

	return grpcOptions, err
}
