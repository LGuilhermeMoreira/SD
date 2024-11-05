package main

import (
	"chat_p2p/server/handler"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		socket, err := list.Accept()
		if err != nil {
			panic(err)
		}
		go handler.HandleConnection(socket)
	}
}
