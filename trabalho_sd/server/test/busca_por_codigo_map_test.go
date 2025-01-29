package test

import (
	"encoding/json"
	"net"
	"testing"
	"time"
	"trabalho_sd/dto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBuscarAlunoPorCodigoMap(t *testing.T) {
	// crio o endereço
	addr, err := net.ResolveUDPAddr("udp", "localhost:4567")
	if err != nil {
		t.Fatalf("Erro ao resolver endereço do servidor: %v", err)
	}

	// cria a conexão com servidor
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		t.Fatalf("Erro ao conectar: %v", err)
	}
	defer conn.Close()

	codigo := map[string]any{
		"codigo": "SD001",
	}

	data, err := json.Marshal(codigo)
	if err != nil {
		t.Fatalf("Erro ao serializar 'codigo': %v", err)
	}

	// montando a mensagem.
	message := dto.Message{
		RequestID:       uuid.New(),
		MessageType:     0,
		ObjectReference: "Escola",
		Arguments:       data,
		Method:          "BuscarAlunoPorCodigo",
	}

	dataR, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Erro ao serializar 'message': %v", err)
	}

	_, err = conn.Write(dataR)
	if err != nil {
		t.Fatalf("Erro ao enviar mensagem: %v", err)
	}

	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	n, _, err := conn.ReadFromUDP(buffer)

	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			t.Fatalf("Timeout ao receber resposta: %v", err)
		} else {
			t.Fatalf("Erro ao receber resposta: %v", err)
		}
	}

	var response dto.Message
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		t.Fatalf("Erro ao deserializar resposta: %v", err)
	}
	assert.NotEmpty(t, response)
}
