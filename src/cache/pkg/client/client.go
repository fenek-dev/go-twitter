package geo

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCacheGrpcClient(url string) (proto.CacheServiceClient, error) {
	var opts []grpc.DialOption = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(url, opts...)
	if err != nil {
		return nil, err
	}

	return proto.NewCacheServiceClient(conn), nil
}
