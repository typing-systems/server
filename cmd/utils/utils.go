package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/typing-systems/typing-server/cmd/pubsub"
)

type cfgStruct struct {
	Debug bool
}

var Config cfgStruct

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
	if Config.Debug {
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

func LoadConfig() {
	if _, err := os.Stat("config.json"); errors.Is(err, os.ErrNotExist) {
		// config.json does not exist
		Config = genConfig()
	} else if err != nil {
		log.Fatalf("error detecting if config file exists: %v", err)
	} else {
		f, err := os.Open("config.json")
		if err != nil {
			log.Fatalf("error opening config file: %v", err)
		}
		decoder := json.NewDecoder(f)
		cfg := cfgStruct{}
		err = decoder.Decode(&cfg)
		if err != nil {
			log.Fatalf("error decoding json from config file: %v", err)
		}
		Config = cfg
		if err := f.Close(); err != nil {
			log.Fatalf("error closing config file: %v", err)
		}
	}
}

func genConfig() cfgStruct {
	// default configuration
	cfg := cfgStruct{
		Debug: false,
	}

	jsonCfg, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		log.Fatalf("error generating config file json: %v", err)
	}

	// 0600 file perm means read/write by owner only
	err = ioutil.WriteFile("config.json", jsonCfg, 0600)
	if err != nil {
		log.Fatalf("error writing config file: %v", err)
	}

	return cfg
}
