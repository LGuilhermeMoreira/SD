package handlers

import "net"

func HandleRead(conn net.Conn) {
	defer conn.Close()
	conn.Read()
}
