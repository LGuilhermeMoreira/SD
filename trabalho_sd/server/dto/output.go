package dto

import "trabalho_sd/models"

type OutputBuscaAlunoPorCodigo struct {
	Alunos []models.Aluno `json:"alunos"`
	Status int            `json:"status"`
}

type OutPutStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
