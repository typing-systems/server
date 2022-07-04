package connections

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/typing-systems/typing-server/cmd/lobby"
)

func HandleConnection(c net.Conn) {
	log.Println("New connection from", c.RemoteAddr().String())

connection:
	for {
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
		}

		cleanedMessage := strings.TrimSpace(message)
		switch cleanedMessage {
		case "newConnection":
			lobbyID, clients := lobby.NewConnection(c.RemoteAddr().String())
			c.Write([]byte("connected to: " + lobbyID + "\ncurrent clients: " + strconv.Itoa(clients)))

		case "close":
			c.Write([]byte("bye!\n"))
			break connection

		case "ping":
			c.Write([]byte("pong\n"))

		case "pong":
			c.Write([]byte("ping\n"))
		}
	}
	c.Close()
}
