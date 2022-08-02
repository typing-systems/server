package connections

import (
	"context"
	"fmt"
	"log"

	"github.com/typing-systems/typing-server/cmd/db"
	"github.com/typing-systems/typing-server/cmd/utils"
)

type Server struct{}

var b = utils.InstantiateBroker()

func (s *Server) Connected(ctx context.Context, e *Empty) (*MyPosition, error) {
	log.Printf("\033[32mClient connected\033[0m")

	uuid := utils.GenerateUUID()
	lobbyID, lane, isNewLobby := utils.Matchmake(uuid, b.GetAllLobbies())

	if isNewLobby {
		fmt.Printf("\033[33mMaking a new lobby: %s\033[0m\n", uuid)
		b.AddLobby(lobbyID)
	}

	return &MyPosition{LobbyID: lobbyID, Lane: lane}, nil
}

func (s *Server) UpdatePosition(ctx context.Context, myPosition *MyPosition) (*Empty, error) {
	db.UpdatePosition(myPosition.LobbyID, myPosition.Lane)

	fmt.Println("calling publish on", myPosition.LobbyID, "for lane", myPosition.Lane)
	newPoint := b.IncrPoints(myPosition.LobbyID, myPosition.Lane)
	b.Publish(myPosition.LobbyID, myPosition.Lane, newPoint)
	fmt.Println("publish called")

	return &Empty{}, nil
}

func (s *Server) Positions(lobby *MyLobby, stream Connections_PositionsServer) error {
	fmt.Println("positions called")

	data := b.GetLobby(lobby.LobbyID).GetDataChan()

	for {
		if data, ok := <-data; ok {
			fmt.Printf("Lobby %s, received a fucking update: for lane: %s, newPoints: %d\n", lobby.LobbyID, data.GetLane(), data.GetPoints())
			msg := &NewPosition{Lane: data.GetLane(), Points: data.GetPoints32()}
			err := stream.Send(msg)
			if err != nil {
				log.Fatalf("error sending to stream: %v", err)
			}
			fmt.Println("stream sent")
		}
	}
}
