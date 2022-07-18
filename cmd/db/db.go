package db

import (
	"log"

	"github.com/go-redis/redis"
)

var db = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func UpdatePosition(key string, fld string) {
	db.HIncrBy(key, fld, 1)
}

func GetPositionInfo(key string) ([]string, error) {
	results, err := db.HMGet(key, "lane1", "lane2", "lane3", "lane4").Result()
	if err != nil {
		return nil, err
	}

	log.Printf("results: %s | lobby: %s", results, key)

	asserted := make([]string, 4)
	for i, lane := range results {
		if lane == nil {
			log.Fatalf("lane %d is nil!", i)
		}
		asserted[i] = lane.(string)
	}

	return asserted, nil
}

func InitLobby(key string) {
	log.Printf("initlobby: %s", key)
	db.HSet(key, "playerCount", "1")
	db.HSet(key, "lane1", "0")
	db.HSet(key, "lane2", "0")
	db.HSet(key, "lane3", "0")
	db.HSet(key, "lane4", "0")
}

func GetAllLobbies() []string {
	keys, err := db.Keys("*").Result()
	if err != nil {
		log.Fatalf("error retrieving all lobbies: %v", err)
	}

	return keys
}

func GetPlayerCount(lobby string) string {
	playerCount, err := db.HGet(lobby, "playerCount").Result()
	if err != nil {
		log.Fatalf("error getting player count: %v", err)
	}
	return playerCount
}

func IncrPlayerCount(lobby string) {
	db.HIncrBy(lobby, "playerCount", 1)
}
