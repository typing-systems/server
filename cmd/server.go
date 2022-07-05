package main

import (
	"log"
	"net"

	"github.com/typing-systems/typing-server/cmd/chat"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer()

	s := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("failed to serve", err)
	}
}
