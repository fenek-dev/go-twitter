package geo

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
}

func New(url string) (*Client, error) {
	var opts []grpc.DialOption = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(url, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil

}

func (c *Client) NewService() proto.AuthServiceClient {
	return proto.NewAuthServiceClient(c.conn)
}

func (c *Client) Close() error {
	return c.conn.Close()
}
