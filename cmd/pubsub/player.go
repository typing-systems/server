package pubsub

import (
	"sync"
)

type Player struct {
	id     string
	active bool
	mutex  sync.RWMutex
}

func CreateNewPlayer(uuid string) *Player {
	return &Player{
		id:     uuid,
		active: true,
	}
}
