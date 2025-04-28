package sdk

import (
	"fmt"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/nix_sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IClient interface {
	pb.NixClient
}

type Client struct {
	option *Options

	conn *grpc.ClientConn

	pb.NixClient
}

var _ pb.NixClient = (*Client)(nil)
var _ IClient = (*Client)(nil)

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
		log.Logger.Error(fmt.Sprintf("occur error when create nix grpc server connection , host:%s,error: %s\n",
			c.option.getHostTarget(),
			err.Error()))
		return err
	}
	log.Logger.Info(fmt.Sprintf("initialize nix sdk grpc connection finished,host:%s\n", c.option.getHostTarget()))
	c.conn = conn
	//保存客户端
	c.NixClient = pb.NewNixClient(conn)

	return nil
}
