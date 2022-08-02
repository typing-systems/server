package utils

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/typing-systems/typing-server/cmd/pubsub"
)

func GenerateUUID() string {
	id, err := uuid.NewV4()

	if err != nil {
		log.Fatalf("Error with UUID: %v", err)
	}

	return id.String()
}

func Matchmake(uuid string, lobbies pubsub.Lobbies) (string, string, bool) {
	for _, lobby := range lobbies {
		playerCount := lobby.GetPlayerCount()
		fmt.Printf("playerCount of %s is: %d\n", lobby.GetLobbyID(), lobby.GetPlayerCount())
		if playerCount < 4 {
			lobby.IncrPlayerCount()
			fmt.Printf("adding client to %s\n", lobby.GetLobbyID())
			return lobby.GetLobbyID(), "lane" + strconv.Itoa(playerCount+1), false
		}
	}

	return uuid, "lane1", true
}

func InstantiateBroker() *pubsub.Broker {
	return pubsub.NewBroker()
}
