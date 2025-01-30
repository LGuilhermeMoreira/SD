package udp

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"trabalho_sd/dto"

	"github.com/google/uuid"
)

type Dispatcher struct {
	Skeleton  Skeleton
	mensagens map[uuid.UUID]bool
}

func NewDispatcher(s Skeleton) *Dispatcher {
	return &Dispatcher{
		Skeleton:  s,
		mensagens: make(map[uuid.UUID]bool),
	}
}

func (d *Dispatcher) Solve(conn *net.UDPConn, addr *net.UDPAddr, buffer []byte) {
	var msg dto.Message
	json.Unmarshal(buffer, &msg)
	msg.Debug()
	if _, loaded := d.mensagens[msg.RequestID]; loaded {
		msg.Error = map[string]any{
			"status": 409,
			"error":  "Mensagem duplicada",
		}
		msg.MessageType = 1
		data, _ := json.Marshal(msg)
		conn.WriteToUDP(data, addr)
	}
	switch msg.ObjectReference {
	case "Escola":
		d.Skeleton.HandleRequest(&msg)
	default:
		msg.MessageType = 1
		msg.Error = map[string]any{
			"status": 404,
			"error":  "objectReference não encontrado",
		}
	}
	msg.Debug()
	randNum := rand.Int()
	if randNum%2 == 0 && randNum%3 == 0 && randNum%5 == 0 {
		log.Println("Erro ao enviar a mensagem", randNum)
		return
	}
	data, _ := json.Marshal(msg)
	conn.WriteToUDP(data, addr)
}
