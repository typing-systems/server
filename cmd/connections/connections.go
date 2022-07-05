package connections

import (
	"context"
	"errors"
	"log"

	"github.com/typing-systems/typing-server/cmd/db"
	"github.com/typing-systems/typing-server/cmd/lobby"
)

type Server struct{}

func (s *Server) Connected(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Client connected, ID: %s", message.Body)

	lobbyID, err := lobby.Matchmake()
	if err != nil {
		return nil, errors.New("Error finding a game")
	}
	db.PlayerHSet(message.Body, "lobbyID", lobbyID)
	log.Printf("player lobby: %s", db.PlayerHGet(message.Body, "lobbyID"))
	return &Message{Body: "Game found!"}, nil
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello from the Server!"}, nil
}
