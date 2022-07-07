package lobby

import (
	"log"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/typing-systems/typing-server/cmd/db"
)

func Matchmake() (string, string) {
	for _, lobby := range db.GetAllLobbies() {
		log.Printf("lobby: %s", lobby)
		playerCount, err := strconv.Atoi(db.GetPlayerCount(lobby))
		if err != nil {
			log.Fatalf("error converting playerCount to int: %v", err)
		}

		if playerCount < 4 {
			db.IncrPlayerCount(lobby)
			return lobby, ("lane" + strconv.Itoa(playerCount+1))
		}
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Error with UUID: %v", err)
	}
	idString := id.String()

	db.InitLobby(idString)

	return idString, "lane1"
}
