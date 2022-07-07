package connections

import (
	"context"
	"log"

	"github.com/typing-systems/typing-server/cmd/db"
	"github.com/typing-systems/typing-server/cmd/lobby"
)

type Server struct{}

func (s *Server) Connected(ctx context.Context, e *Empty) (*MyPosition, error) {
	log.Printf("Client connected")

	lobbyID, lane := lobby.Matchmake()
	db.InitLobby(lobbyID)
	return &MyPosition{ID: lobbyID, Lane: lane}, nil
}

func (s *Server) Positions(ctx context.Context, myPosition *MyPosition) (*PositionInfo, error) {
	positions := db.UpdatePosition(myPosition.ID, myPosition.Lane)

	return &PositionInfo{Lane1: positions[0], Lane2: positions[1], Lane3: positions[2], Lane4: positions[3]}, nil
}
