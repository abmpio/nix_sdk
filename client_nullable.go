package nix_sdk

import (
	"context"

	pb "github.com/abmpio/nix_sdk/proto"
	"google.golang.org/grpc"
)

type NullableClient struct {
}

var _ pb.NixClient = (*NullableClient)(nil)

func (c *NullableClient) HealthCheck(ctx context.Context, in *pb.NixHealthCheckRequest, opts ...grpc.CallOption) (*pb.NixHealthCheckResponse, error) {
	return &pb.NixHealthCheckResponse{
		Status: pb.NixHealthCheckResponse_NOT_SERVING,
	}, nil
}
func (c *NullableClient) FindOneKVByKey(ctx context.Context, in *pb.FindKVOneByKeyRequest, opts ...grpc.CallOption) (*pb.FindKVOneByKeyResponse, error) {
	return &pb.FindKVOneByKeyResponse{}, nil
}

func (c *NullableClient) FindKVListByKeyList(ctx context.Context, in *pb.FindKVListByKeyListRequest, opts ...grpc.CallOption) (*pb.FindKVListByKeyListResponse, error) {
	return &pb.FindKVListByKeyListResponse{}, nil
}

func (c *NullableClient) FindKVListByTag(ctx context.Context, in *pb.FindKVListByTagRequest, opts ...grpc.CallOption) (*pb.FindKVListByTagResponse, error) {
	return &pb.FindKVListByTagResponse{}, nil
}
