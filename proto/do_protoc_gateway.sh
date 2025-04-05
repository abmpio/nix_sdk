#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"
mkdir ./gateway/ -p
protoc -I . --grpc-gateway_out ./gateway/ \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt grpc_api_configuration=./openapi_service.yaml \
    --grpc-gateway_opt standalone=true \
    ./nix.proto