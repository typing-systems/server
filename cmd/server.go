package main

import (
	"log"
	"net"

	"github.com/typing-systems/typing-server/cmd/connections"
	"github.com/typing-systems/typing-server/cmd/settings"
	"google.golang.org/grpc"
)

func main() {
	settings.Load("prod")

	l, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer()

	s := connections.Server{}
	connections.RegisterConnectionsServer(grpcServer, &s)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("failed to serve", err)
	}
}
