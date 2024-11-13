package handler

import (
	"encoding/json"
	"net"
	"sockets/entity"
	"sockets/utils"
)

func Connection(conn net.Conn) {
	var req entity.Request
	err := json.NewDecoder(conn).Decode(&req)
	if err != nil {
		json.NewEncoder(conn).Encode(entity.Response{
			Result: err.Error(),
		})
		return
	}
	var resp entity.Response
	switch req.Operation {
	case "**":
		resp = utils.GetReponse(entity.ScienceCalculator{}, req)
	case "@":
		resp = utils.GetReponse(entity.AnotherCalculator{}, req)
	default:
		resp = utils.GetReponse(entity.SimpleCalculator{}, req)
	}
	json.NewEncoder(conn).Encode(resp)
	return
}
