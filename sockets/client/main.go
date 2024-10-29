package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	port := flag.String("port", "8080", "The port to connect to")
	msg := flag.String("msg", "hello", "The message to send")
	conn, err := net.Dial("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte(*msg))
}
