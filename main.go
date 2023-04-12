package main

import (
	"context"
	"github.com/blackmenthor/grpc_sample/tutorial"
	rkboot "github.com/rookie-ninja/rk-boot"
	rkgrpc "github.com/rookie-ninja/rk-grpc/boot"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	println("gRPC server tutorial in Go")

	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Get grpc entry with name
	grpcEntry := boot.GetEntry("tutorial").(*rkgrpc.GrpcEntry)
	grpcEntry.AddRegFuncGrpc(registerTutorialService)

	// *************************************** //
	// *** Set server receive size to 20MB ***
	// *************************************** //
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	grpcEntry.AddServerOptions(grpc.MaxRecvMsgSize(1 * gb))
	grpcEntry.AddServerOptions(grpc.MaxSendMsgSize(1 * gb))
	grpcEntry.AddServerOptions(grpc.MaxMsgSize(1 * gb))

	// Bootstrap
	boot.Bootstrap(context.Background())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())

	//listener, err := net.Listen("tcp", ":9000")
	//if err != nil {
	//	panic(err)
	//}
	//
	//s := grpc.NewServer(
	//	grpc.MaxRecvMsgSize(10010241024),
	//	grpc.MaxSendMsgSize(10010241024),
	//	grpc.MaxHeaderListSize(uint32(1000000000)),
	//)
	//tutorial.RegisterAlbumServiceServer(s, &tutorial.Server{})
	//if err := s.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}

func registerTutorialService(s *grpc.Server) {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	tutorial.RegisterAlbumServiceServer(s, &tutorial.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
