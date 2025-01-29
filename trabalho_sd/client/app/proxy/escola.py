from app.models import Aluno, Codigo, Disciplina, Professor, Message
from app.connection import enviar_mensagem

class EscolaService:
    
    def cadastrar_aluno(self, aluno: Aluno):
       msg = Message("Escola","CadastrarAluno",aluno)
       response = enviar_mensagem(message=msg)
       print(response)
    

    def cadastrar_disciplina(self, disciplina: Disciplina):
       msg = Message("Escola","CadastrarDisciplina",disciplina)
       response = enviar_mensagem(message=msg)
       print(response)
        

  
    def cadastrar_professor(self, professor: Professor):
        msg = Message("Escola","CadastrarProfessor",professor)
        response = enviar_mensagem(message=msg)
        print(response)
    
   
    def buscar_aluno_por_codigo(self, codigo: Codigo):
        msg = Message("Escola","BuscarAlunoPorCodigo",codigo)
        response = enviar_mensagem(message=msg)
        print(response)