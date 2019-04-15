package main

import (
	"net"

	pb "github.com/mas9612/authserver/pkg/authserver"
	"github.com/mas9612/authserver/pkg/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		logger.Fatal("listen failed", zap.Error(err))
	}
	defer listener.Close()

	s, err := server.NewAuthserver(logger, server.SetAddr("10.1.3.21"))
	if err != nil {
		logger.Fatal("failed to initialize server", zap.Error(err))
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthserverServer(grpcServer, s)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("service failed", zap.Error(err))
	}
}
