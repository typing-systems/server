package lobby

type lobby struct {
	lobbyID string
	player1 string
	player2 string
	player3 string
	player4 string

	clients int
}

var lobbyMap = make(map[string]int)
var lobbyList []lobby

func NewConnection(clientID string) (string, int) {
	if _, alreadyConnected := lobbyMap[clientID]; alreadyConnected {
		return "client already connected to " + lobbyList[lobbyMap[clientID]].lobbyID, lobbyList[lobbyMap[clientID]].clients
	}

	for i, lobby := range lobbyList {
		if lobby.clients < 4 {
			switch lobby.clients {
			case 1:
				lobby.player2 = clientID
				lobby.clients++
				lobbyMap[clientID] = i
				return lobby.lobbyID, lobby.clients
			case 2:
				lobby.player3 = clientID
				lobby.clients++
				lobbyMap[clientID] = i
				return lobby.lobbyID, lobby.clients
			case 3:
				lobby.player4 = clientID
				lobby.clients++
				lobbyMap[clientID] = i
				return lobby.lobbyID, lobby.clients
			}
		}
	}
	lobbyMap[clientID] = len(lobbyList)
	lobbyList = append(lobbyList, lobby{
		lobbyID: "lobby" + clientID,
		player1: clientID,
		clients: 1,
	})
	return lobbyList[lobbyMap[clientID]].lobbyID, lobbyList[lobbyMap[clientID]].clients
}
