package connections

import (
	"context"
	"errors"
	"log"

	"github.com/typing-systems/typing-server/cmd/db"
	"github.com/typing-systems/typing-server/cmd/lobby"
)

var errGameNotFound = errors.New("error game not found")

type Server struct{}

func (s *Server) Connected(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Client connected, ID: %s", message.Body)

	lobbyID, err := lobby.Matchmake()
	if err != nil {
		return nil, errGameNotFound
	}
	db.LobbySetZero(lobbyID)
	db.PlayerHSet(message.Body, "lobbyID", lobbyID)
	log.Printf("player lobby: %s", db.PlayerHGet(message.Body, "lobbyID"))
	return &Message{Body: "Game found!"}, nil
}

func (s *Server) Positions(ctx context.Context, myPosition *MyPosition) (*PositionInfo, error) {
	positions := db.LobbyUpdatePosition("lobbyID", myPosition.Lane)

	return &PositionInfo{Lane1: positions[0], Lane2: positions[1], Lane3: positions[2], Lane4: positions[3]}, nil
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello from the Server!"}, nil
}
