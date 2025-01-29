package db

import (
	"trabalho_sd/models"
)

type BancoDeDados struct {
	Alunos      []models.Aluno
	Disciplinas []models.Disciplina
	Professores []models.Professor
}

func NewBancoDeDados() *BancoDeDados {
	return &BancoDeDados{
		Professores: []models.Professor{
			*models.NewProfessor("samy", "samy@mail.com"),
			*models.NewProfessor("marcos", "marcos@email.com"),
			*models.NewProfessor("thigas", "thigas@email.com"),
			*models.NewProfessor("livia", "livia@email.com"),
		},
		Disciplinas: []models.Disciplina{
			*models.NewDisciplina("SD", "SD001", *models.NewProfessor("marcos", "marcos@email.com")),
			*models.NewDisciplina("SO", "SO001", *models.NewProfessor("thigas", "thigas@email.com")),
			*models.NewDisciplina("BD", "BD001", *models.NewProfessor("livia", "livia@email.com")),
			*models.NewDisciplina("IA", "IA001", *models.NewProfessor("samy", "samy@mail.com")),
		},
		Alunos: []models.Aluno{
			*models.NewAluno("Guigui", "CC", "1233212221",
				*models.NewDisciplina("SD", "SD001",
					*models.NewProfessor("marcos", "marcos@email.com"))),

			*models.NewAluno("Iaia Pirata", "CC", "1233212221",
				*models.NewDisciplina("SD", "SD001",
					*models.NewProfessor("marcos", "marcos@email.com"))),

			*models.NewAluno("Kaynara", "CC", "1233212222",
				*models.NewDisciplina("SO", "SO001",
					*models.NewProfessor("thigas", "thigas@email.com"))),

			*models.NewAluno("Cabeça", "CC", "1233212223",
				*models.NewDisciplina("BD", "BD001",
					*models.NewProfessor("livia", "livia@email.com"))),

			*models.NewAluno("jss de CC", "CC", "1233212224",
				*models.NewDisciplina("IA", "IA001",
					*models.NewProfessor("samy", "samy@mail.com"))),
		},
	}
}

func (b *BancoDeDados) CadastrarAluno(aluno models.Aluno) error {
	b.Alunos = append(b.Alunos, aluno)
	return nil
}

func (b *BancoDeDados) CadastrarDisciplina(disciplina models.Disciplina) error {
	b.Disciplinas = append(b.Disciplinas, disciplina)
	return nil
}

func (b *BancoDeDados) CadastrarProfessor(professor models.Professor) error {
	b.Professores = append(b.Professores, professor)
	return nil
}

func (b *BancoDeDados) BuscarAlunoPorCodigo(codigo string) ([]models.Aluno, error) {
	var responseData []models.Aluno
	for _, aluno := range b.Alunos {
		for _, disciplina := range aluno.Disciplinas {
			if disciplina.Codigo == codigo {
				responseData = append(responseData, aluno)
			}
		}
	}
	return responseData, nil
}
