package pubsub

import (
	"sync"
)

type Lobby struct {
	id          string
	exists      bool
	playerCount int
	lanes       []int
	data        chan *Data
	mutex       sync.RWMutex
}

func CreateNewLobby(uuid string) *Lobby {
	newLobby := Lobby{
		id:          uuid,
		exists:      true,
		playerCount: 1,
		lanes:       []int{0, 0, 0, 0},
		data:        make(chan *Data),
	}
	return &newLobby
}

func (l *Lobby) IncrPlayerCount() {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	l.playerCount++
}

func (l *Lobby) DecrPlayerCount() {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	l.playerCount--
}

func (l *Lobby) GetPlayerCount() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.playerCount
}

func (l *Lobby) GetDataChan() chan *Data {
	return l.data
}

func (l *Lobby) GetLanes() []int {
	return l.lanes
}

func (l *Lobby) GetLobbyID() string {
	return l.id
}

func (l *Lobby) Destruct() {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	l.exists = false
	close(l.data)
}

func (l *Lobby) Signal(data *Data) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if l.exists {
		l.data <- data
	}
}
