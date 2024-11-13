package entity

import (
	"encoding/json"
	"fmt"
	"net"
	"sockets/utils"
)

const port = ":8080"

type User struct {
	input  string
	Result Response
}

func NewUser(input string) *User {
	return &User{
		input: input,
	}
}

func (u *User) SendRequest() error {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		return err
	}
	defer conn.Close()
	req, err := utils.InputToRequest(u.input)
	if err != nil {
		return err
	}
	res, err := handleWriteRead(conn, *req)
	if err != nil {
		return err
	}
	u.Result = *res
	return nil
}

func (u *User) ShowResponse() {
	fmt.Println(u.Result.Result)
}

func handleWriteRead(conn net.Conn, req Request) (*Response, error) {
	err := json.NewEncoder(conn).Encode(req)
	if err != nil {
		return nil, err
	}
	var res Response
	err = json.NewDecoder(conn).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
