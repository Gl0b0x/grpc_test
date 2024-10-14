package main

import (
	desc "github.com/Gl0b0x/grpc_test/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}
