package udp

import (
	"encoding/json"
	"trabalho_sd/dto"
	"trabalho_sd/models"
	"trabalho_sd/service"
)

type Skeleton struct {
	escolaService service.Escola
}

func NewSkeleton(escolaService service.Escola) *Skeleton {
	return &Skeleton{
		escolaService: escolaService,
	}
}

func (s *Skeleton) HandleRequest(message *dto.Message) {
	switch message.Method {
	case "CadastrarAluno":
		var alunoDto dto.Aluno
		json.Unmarshal(message.Arguments, &alunoDto)
		data, ok := s.handleCadastarAluno(alunoDto)
		escreveMensagem(data, ok, message)
	case "CadastrarProfessor":
		var professorDto dto.Professor
		json.Unmarshal(message.Arguments, &professorDto)
		data, ok := s.handleCadastarProfessor(professorDto)
		escreveMensagem(data, ok, message)
	case "CadastrarDisciplina":
		var disciplinaDto dto.Disciplina
		json.Unmarshal(message.Arguments, &disciplinaDto)
		data, ok := s.handleCadastarDisciplina(disciplinaDto)
		escreveMensagem(data, ok, message)
	case "BuscarAlunoPorCodigo":
		var codigoDto dto.Codigo
		json.Unmarshal(message.Arguments, &codigoDto)
		data, ok := s.handleBuscarAlunoPorCodigo(codigoDto.Codigo)
		escreveMensagem(data, ok, message)
	default:
		data := map[string]any{
			"status": 404,
			"error":  "serviço não encontrado",
		}
		ok := false
		escreveMensagem(data, ok, message)
	}
}

func (s *Skeleton) handleCadastarAluno(alunoDto dto.Aluno) (any, bool) {
	alunoModel := models.NewAluno(alunoDto.Nome, alunoDto.Curso, alunoDto.Cpf, alunoDto.Disciplinas...)
	response, err := s.escolaService.CadastrarAluno(*alunoModel)
	if err != nil {
		return map[string]any{
			"status": 500,
			"error":  err.Error(),
		}, false
	}
	return response, true
}

func (s *Skeleton) handleCadastarProfessor(professorDto dto.Professor) (any, bool) {
	professorModel := models.NewProfessor(professorDto.Nome, professorDto.Email)
	response, err := s.escolaService.CadastrarProfessor(*professorModel)
	if err != nil {
		return map[string]any{
			"status": 500,
			"error":  err.Error(),
		}, false
	}
	return response, true
}

func (s *Skeleton) handleCadastarDisciplina(disciplinaDto dto.Disciplina) (any, bool) {
	disciplinaModel := models.NewDisciplina(disciplinaDto.Nome, disciplinaDto.Codigo, disciplinaDto.Professores...)
	response, err := s.escolaService.CadastrarDisciplina(*disciplinaModel)
	if err != nil {
		return map[string]any{
			"status": 500,
			"error":  err.Error(),
		}, false
	}
	return response, true
}

func (s *Skeleton) handleBuscarAlunoPorCodigo(codigo string) (any, bool) {
	response, err := s.escolaService.BuscarAlunoPorCodigo(codigo)
	if err != nil {
		return map[string]any{
			"status": 500,
			"error":  err.Error(),
		}, false
	}
	return response, true
}

func escreveMensagem(data any, ok bool, message *dto.Message) {
	if !ok {
		message.Error = data
	} else {
		responseData, _ := json.Marshal(data)
		message.Arguments = responseData
	}
	message.MessageType = 1
}
