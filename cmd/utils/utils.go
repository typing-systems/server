package utils

import (
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

func Matchmake(uuid string, lobbies pubsub.Lobbies) (string, bool) {
	for _, lobby := range lobbies {
		playerCount := lobby.GetPlayerCount()

		if playerCount < 4 {
			lobby.IncrPlayerCount()
			return "lane" + strconv.Itoa(playerCount+1), false
		}
	}

	return "lane1", true
}

func InstantiateBroker() *pubsub.Broker {
	return pubsub.NewBroker()
}
