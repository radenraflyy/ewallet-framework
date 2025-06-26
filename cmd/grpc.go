package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	listener, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("Failed to listen grpc port: ", err)
	}

	s := grpc.NewServer()

	// list method
	logrus.Info("gRPC server is running on port: ", helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to serve grpc: ", err)
	}
}
