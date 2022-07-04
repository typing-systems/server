package main

import (
	"log"
	"net"

	"github.com/typing-systems/typing-server/cmd/connections"
)

func main() {
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go connections.HandleConnection(c)
	}
}
