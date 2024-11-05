package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		var input string
		fmt.Scanln(&input)
		conn.Write([]byte(input))
		buff := make([]byte, 1024)
		_, err := conn.Read(buff)
		if err != nil {
			panic(err)
		}
		log.Println(string(buff))
	}
}
