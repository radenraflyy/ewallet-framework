package cmd

import (
	"log"
	"net"
	"ums/helpers"

	pb "ums/cmd/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	// init dependencies
	depedency := depedencyInject()

	s := grpc.NewServer()

	// list method
	pb.RegisterTokenValidationServiceServer(s, depedency.TokenValidationApi)

	listener, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("Failed to listen grpc port: ", err)
	}

	logrus.Info("gRPC server is running on port: ", helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to serve grpc: ", err)
	}
}
