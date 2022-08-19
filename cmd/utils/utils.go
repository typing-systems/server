package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/typing-systems/typing-server/cmd/pubsub"
	"github.com/typing-systems/typing-server/cmd/settings"
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
			fmt.Printf("\033[33madding client to %s\033[0m\n", lobby.GetLobbyID())
			return lobby.GetLobbyID(), "lane" + strconv.Itoa(playerCount+1), false
		}
	}

	return uuid, "lane1", true
}

func InstantiateBroker() *pubsub.Broker {
	return pubsub.NewBroker()
}

func Log(text string) {
	if settings.Values.Debug {
		f, err := os.OpenFile("./server.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		if _, err = f.WriteString(time.Now().Format("01-02-2006 15:04:05.000000		") + text + "\n"); err != nil {
			panic(err)
		}

		if err := f.Close(); err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}
}
