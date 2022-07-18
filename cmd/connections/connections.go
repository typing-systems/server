package connections

import (
	"context"
	"log"
	"time"

	"github.com/typing-systems/typing-server/cmd/db"
	"github.com/typing-systems/typing-server/cmd/lobby"
)

type Server struct{}

func (s *Server) Connected(ctx context.Context, e *Empty) (*MyPosition, error) {
	log.Printf("Client connected")

	lobbyID, lane := lobby.Matchmake()
	db.InitLobby(lobbyID)
	return &MyPosition{LobbyID: lobbyID, Lane: lane}, nil
}

func (s *Server) UpdatePosition(ctx context.Context, myPosition *MyPosition) (*Empty, error) {
	db.UpdatePosition(myPosition.LobbyID, myPosition.Lane)
	return &Empty{}, nil
}

func (s *Server) Positions(lobbyID *MyLobby, stream Connections_PositionsServer) error {
	for {
		positionInfo, err := db.GetPositionInfo(lobbyID.LobbyID)
		if err != nil {
			log.Fatalf("error calling db.GetPositionInfo: %v", err)
		}
		stream.Send(&PositionInfo{Lane1: positionInfo[0], Lane2: positionInfo[1], Lane3: positionInfo[2], Lane4: positionInfo[3]})
		time.Sleep(500 * time.Millisecond)
	}
}
