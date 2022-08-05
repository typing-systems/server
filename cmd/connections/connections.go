package connections

import (
	"context"
	"log"
	"strconv"

	"github.com/typing-systems/typing-server/cmd/db"
	"github.com/typing-systems/typing-server/cmd/utils"
)

type Server struct{}

var b = utils.InstantiateBroker()

func (s *Server) Connected(ctx context.Context, e *Empty) (*MyPosition, error) {
	utils.Log("New client connected")

	uuid := utils.GenerateUUID()
	lobbyID, lane, isNewLobby := utils.Matchmake(uuid, b.GetAllLobbies())

	if isNewLobby {
		utils.Log("Making a new lobby: " + uuid)
		b.AddLobby(lobbyID)
	}

	return &MyPosition{LobbyID: lobbyID, Lane: lane}, nil
}

func (s *Server) UpdatePosition(ctx context.Context, myPosition *MyPosition) (*Empty, error) {
	db.UpdatePosition(myPosition.LobbyID, myPosition.Lane)

	utils.Log("Calling UpdatePosition on lobby" + myPosition.LobbyID + " for " + myPosition.Lane)
	newPoint := b.IncrPoints(myPosition.LobbyID, myPosition.Lane)
	b.Publish(myPosition.LobbyID, myPosition.Lane, newPoint)
	utils.Log("UpdatePosition finished")

	return &Empty{}, nil
}

func (s *Server) Positions(lobby *MyLobby, stream Connections_PositionsServer) error {
	utils.Log("Positions called")

	data := b.GetLobby(lobby.LobbyID).GetDataChan()

	for {
		if data, ok := <-data; ok {
			utils.Log("Sending lane " + data.Lane + " and points " + strconv.Itoa(data.GetPoints()) + "to " + lobby.LobbyID)
			msg := &NewPosition{Lane: data.GetLane(), Points: data.GetPoints32()}
			err := stream.Send(msg)
			if err != nil {
				log.Fatalf("error sending to stream: %v", err)
			}
			utils.Log("Stream sent")
		}
	}
}
