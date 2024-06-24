package geo

import (
	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewSsoGrpcClient(url string) (ssov1.AuthServiceClient, error) {
	var opts []grpc.DialOption = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(url, opts...)
	if err != nil {
		return nil, err
	}

	return ssov1.NewAuthServiceClient(conn), nil
}
