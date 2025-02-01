package nix_sdk

import (
	"log"

	pb "github.com/abmpio/nix_sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	option *Options

	conn *grpc.ClientConn

	pb.NixClient
}

var _ pb.NixClient = (*Client)(nil)

func NewClient(opts ...Option) *Client {
	client := &Client{
		option: newDefaultOptions(),
	}
	for _, eachOpt := range opts {
		eachOpt(client.option)
	}
	return client
}

func (c *Client) GetOptions() *Options {
	return c.option
}

// 初始化client
func (c *Client) InitConnnection(opts ...grpc.DialOption) error {
	mergedOpts := make([]grpc.DialOption, 0)
	mergedOpts = append(mergedOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	mergedOpts = append(mergedOpts, grpc.WithKeepaliveParams(*c.option.keepaliveParam))
	mergedOpts = append(mergedOpts, opts...)
	conn, err := grpc.NewClient(c.option.getHostTarget(), mergedOpts...)
	if err != nil {
		log.Printf("occur error when create grpc server connection , host:%s,error: %s\n",
			c.option.getHostTarget(), err.Error())
		return err
	}
	log.Printf("initialize nix_client grpc connection finished,host:%s\n", c.option.getHostTarget())
	c.conn = conn
	//保存客户端
	c.NixClient = pb.NewNixClient(conn)

	return nil
}
