package main

import (
	"flag"
	"log"
	"sockets/entity"
)

const (
	Port = ":8080"
)

func main() {
	input := flag.String("equation", "1 + 1", "equation resolved by calculator")
	flag.Parse()
	user := entity.NewUser(*input)
	if err := user.SendRequest(); err != nil {
		log.Fatal(err)
	}
	user.ShowResponse()
}
