package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := list.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	msg := make([]byte, 1024)
	n, err := conn.Read(msg)
	if err != nil {
		log.Println(conn.RemoteAddr(), err)
	}
	fmt.Println(string(msg[:n]), conn.RemoteAddr())
}
