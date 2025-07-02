package cmd

import (
	"log"
	"net"
	"wallet/helpers"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("Failed to listen grpc port: ", err)
	}

	logrus.Info("gRPC server is running on port: ", helpers.GetEnv("GRPC_PORT", "7000"))
	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to serve grpc: ", err)
	}
}
