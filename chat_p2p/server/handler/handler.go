package handler

import (
	"fmt"
	"log"
	"net"
)

func HandleConnection(socket net.Conn) {
	buffer := make([]byte, 1024)
	_, err := socket.Read(buffer)
	if err != nil {
		log.Println("Error reading from socket")
		return
	}
	log.Println("message", string(buffer))
	var input string
	fmt.Scanln(&input)
	_, err = socket.Write([]byte(input))
	if err != nil {
		log.Println("Error writing to socket")
		return
	}
}
