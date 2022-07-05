package db

import (
	"errors"

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

func PlayerMultiSet(key string, flds []string, vals []string) error {
	if len(flds) != len(vals) {
		return errors.New("db.HSet error: fields and values different lengths")
	}

	for i := range flds {
		playerDB.HSet(key, flds[i], vals[i])
	}
	return nil
}

func PlayerHGet(key string, fld string) string {
	result, err := playerDB.HGet(key, fld).Result()
	if err != nil {
		return "Error retrieving " + key
	}

	return result
}
