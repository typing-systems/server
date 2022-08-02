package connections

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/typing-systems/typing-server/cmd/db"
	"github.com/typing-systems/typing-server/cmd/utils"
)

type Server struct{}

var b = utils.InstantiateBroker()

func (s *Server) Connected(ctx context.Context, e *Empty) (*MyPosition, error) {
	log.Printf("Client connected")

	uuid := utils.GenerateUUID()
	lane, isNewLobby := utils.Matchmake(uuid, b.GetAllLobbies())

	return &MyPosition{LobbyID: uuid, Lane: lane}, nil
}

func (s *Server) UpdatePosition(ctx context.Context, myPosition *MyPosition) (*Empty, error) {
	db.UpdatePosition(myPosition.LobbyID, myPosition.Lane)

	b.Publish(myPosition.LobbyID, myPosition.Lane, 9)

	return &Empty{}, nil
}

func (s *Server) Positions(lobby *MyLobby, stream Connections_PositionsServer) error {
	l := b.GetLobby(lobby.LobbyID)
	data := l.GetDataChan()
	fmt.Println("arite pal-1")
	fmt.Println(data)

	for {
		fmt.Println("arite pal")
		if data, ok := <-data; ok {
			fmt.Printf("Lobby %s, received a fucking update: for lane: %s, newPoints: %d\n", lobby.LobbyID, data.GetLane(), data.GetPoints())
			lanes := l.GetLanes()
			stream.Send(&PositionInfo{Lane1: strconv.Itoa(lanes[0]), Lane2: strconv.Itoa(lanes[1]), Lane3: strconv.Itoa(lanes[2]), Lane4: strconv.Itoa(lanes[3])})
		}
	}
}
