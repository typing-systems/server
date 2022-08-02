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

func (b *Broker) AddLobby(l *Lobby) {
	b.mut.Lock()
	defer b.mut.Unlock()
	b.lobbies[l.GetLobbyID()] = l
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
