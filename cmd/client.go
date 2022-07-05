package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/typing-systems/typing-server/cmd/connections"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := connections.NewConnectionsClient(conn)

	response, err := c.SayHello(context.Background(), &connections.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	gameFound, err := c.Connected(context.Background(), &connections.Message{Body: "myClientID"})
	log.Printf("Response from server: %s", gameFound.Body)
}
