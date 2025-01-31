package test

import (
	"encoding/json"
	"errors"
	"net"
	"testing"
	"time"
	"trabalho_sd/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// 90784ddf-e491-4433-84ab-eb0d973b4c6e uuid
func TestMensagensDuplicadas(t *testing.T) {
	err := send()
	assert.Nil(t, err)
	err = send()
	assert.NotNil(t, err)
}

func send() error {
	addr, err := net.ResolveUDPAddr("udp", "localhost:4567")
	if err != nil {
		return err
	}

	// cria a conexão com servidor
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	codigo := dto.Codigo{
		Codigo: "SD001",
	}

	data, err := json.Marshal(codigo)
	if err != nil {
		return err
	}

	// montando a mensagem.
	id, err := uuid.Parse("90784ddf-e491-4433-84ab-eb0d973b4c6e")
	if err != nil {
		return err
	}
	message := dto.Message{
		RequestID:       id,
		MessageType:     0,
		ObjectReference: "Escola",
		Arguments:       data,
		Method:          "BuscarAlunoPorCodigo",
	}

	dataR, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = conn.Write(dataR)
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	n, _, err := conn.ReadFromUDP(buffer)

	if err != nil {
		return err
	}

	var response dto.Message
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		return err
	}

	if response.Error == nil {
		return nil
	} else {
		return errors.New("erro no campo Error de response")
	}
}
