package pubsub

import (
	"sync"
)

type Lobbies map[string]*Lobby

type Broker struct {
	lobbies Lobbies
	mut     sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		lobbies: Lobbies{},
	}
}

func (b *Broker) AddLobby(uuid string) {
	b.mut.Lock()
	defer b.mut.Unlock()
	l := CreateNewLobby(uuid)
	b.lobbies[uuid] = l
}

func (b *Broker) RemoveLobby(l *Lobby) {
	b.mut.Lock()
	delete(b.lobbies, l.GetLobbyID())
	b.mut.Unlock()
	l.Destruct()
}

func (b *Broker) GetLobby(uuid string) *Lobby {
	return b.lobbies[uuid]
}

func (b *Broker) GetAllLobbies() Lobbies {
	return b.lobbies
}

func (b *Broker) IncrPoints(uuid string, lane string) int {
	l := b.lobbies[uuid]
	switch lane {
	case "lane1":
		l.lanes[0]++
		return l.lanes[0]
	case "lane2":
		l.lanes[1]++
		return l.lanes[1]
	case "lane3":
		l.lanes[2]++
		return l.lanes[2]
	case "lane4":
		l.lanes[3]++
		return l.lanes[3]
	}
	return -1
}

func (b *Broker) Publish(uuid string, lane string, points int) {
	b.mut.RLock()
	defer b.mut.RUnlock()

	d := NewData(lane, points)
	l := b.lobbies[uuid]

	if !l.exists {
		return
	}
	go (func(s *Lobby) {
		l.Signal(d)
	})(l)
}
