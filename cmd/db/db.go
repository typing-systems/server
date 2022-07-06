package db

import (
	"log"

	"github.com/go-redis/redis"
)

var playerDB = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var lobbyDB = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       1,  // use lobby DB
})

func PlayerHSet(key string, fld string, val string) {
	playerDB.HSet(key, fld, val)
}

func PlayerHGet(key string, fld string) string {
	result, err := playerDB.HGet(key, fld).Result()
	if err != nil {
		return "Error retrieving " + key
	}

	return result
}

func LobbyUpdatePosition(key string, fld string) []string {
	lobbyDB.HIncrBy(key, fld, 1)
	results, err := lobbyDB.HMGet(key, "lane1", "lane2", "lane3", "lane4").Result()
	if err != nil {
		return []string{"Error retrieving " + key}
	}

	// log.Printf("r1: %s\nr2: %s\nr3: %s\nr4: %s", r1, r2, r3, r4)
	log.Printf("results: %s", results)

	asserted := make([]string, 4)
	for i, lane := range results {
		if lane == nil {
			log.Fatalf("lane %d is nil!", i)
		}
		asserted[i] = lane.(string)
	}

	return asserted
}

func LobbySetZero(key string) {
	lobbyDB.HSet(key, "lane1", "0")
	lobbyDB.HSet(key, "lane2", "7")
	lobbyDB.HSet(key, "lane3", "6")
	lobbyDB.HSet(key, "lane4", "3")
}
