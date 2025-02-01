package nix_sdk

import (
	pb "github.com/abmpio/nix_sdk/proto"
)

var (
	_globalClient pb.NixClient
)

func GlobalClient() pb.NixClient {
	return _globalClient
}

func SetGlobalClient(c pb.NixClient) {
	_globalClient = c
}
