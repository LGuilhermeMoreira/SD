package udp

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"time"
	"trabalho_sd/dto"

	"github.com/google/uuid"
)

type Dispatcher struct {
	Skeleton  Skeleton
	mensagens map[uuid.UUID]dto.Message
}

func NewDispatcher(s Skeleton) *Dispatcher {
	return &Dispatcher{
		Skeleton:  s,
		mensagens: make(map[uuid.UUID]dto.Message),
	}
}

func (d *Dispatcher) Solve(conn *net.UDPConn, addr *net.UDPAddr, buffer []byte) {
	var msg dto.Message
	err := json.Unmarshal(buffer, &msg)
	if err != nil {
		log.Println("Erro ao deserializar a mensagem:", err)
		return
	}
	msg.Debug()

	// Verifica se a mensagem já foi processada
	if resposta, loaded := d.mensagens[msg.RequestID]; loaded {
		log.Println("Mensagem duplicada detectada:", msg.RequestID)
		// Envia a resposta do histórico (independente do método)
		data, err := json.Marshal(resposta)
		if err != nil {
			log.Println("Erro ao serializar resposta do histórico:", err)
			return
		}
		_, err = conn.WriteToUDP(data, addr)
		if err != nil {
			log.Println("Erro ao enviar resposta do histórico:", err)
		}
		return
	}

	// Processa a mensagem
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

	//Simula perda de pacote aleatoria
	randNum := rand.Int()
	if randNum%2 == 0 && randNum%3 == 0 && randNum%5 == 0 {
		log.Println("Simulando perda de pacote (não enviando a resposta):", randNum)
		return
	}

	// Armazena a resposta no histórico
	d.mensagens[msg.RequestID] = msg

	data, err := json.Marshal(msg)
	if err != nil {
		log.Println("Erro ao serializar a mensagem:", err)
		return
	}
	_, err = conn.WriteToUDP(data, addr)
	if err != nil {
		log.Println("Erro ao enviar a resposta:", err)
	}

	// Remove a mensagem do histórico após um tempo
	go func() {
		time.Sleep(5 * time.Second)
		delete(d.mensagens, msg.RequestID)
		log.Println("Mensagem removida do histórico:", msg.RequestID)
	}()
}
