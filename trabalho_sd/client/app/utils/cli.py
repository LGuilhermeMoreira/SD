from app.models import *
from app.proxy import *
import os
class Interface:
    def __init__(self):
        pass
    
    def start(self):
        es = EscolaService()
        while True:
            print("\nEscolha uma ação:")
            print("1. Buscar Aluno por Código")
            print("2. Cadastrar Aluno")
            print("3. Cadastrar Disciplina")
            print("4. Cadastrar Professor")
            print("0. Sair")

            choice = input("Opção: ")

            if choice == '1':
                codigo_str = input("Digite o código do aluno: ")
                codigo = Codigo(codigo=codigo_str)
                self.limpar_terminal()
                msg = es.buscar_aluno_por_codigo(codigo)
                self.handle_response(msg)
            
            elif choice == '2':
                nome = input("Nome do aluno: ")
                curso = input("Curso do aluno: ")
                cpf = input("CPF do aluno: ")
                prof = Professor("marcos","marcos@email.com").__dict__
                disciplinas = [Disciplina("SD","SD001",prof).__dict__]
                self.limpar_terminal()
                aluno = Aluno(nome=nome, curso=curso, cpf=cpf, disciplinas=disciplinas)
                msg = es.cadastrar_aluno(aluno=aluno)
                self.handle_response(msg)
               

            elif choice == '3':
                nome = input("Nome da disciplina: ")
                codigo = input("Código da disciplina: ")
                professores = []
                self.limpar_terminal()
                disciplina = Disciplina(nome=nome, codigo=codigo, professores=professores)
                msg = es.cadastrar_disciplina(disciplina=disciplina)
                self.handle_response(msg)
            
            elif choice == '4':
                nome = input("Nome do professor: ")
                email = input("Email do professor: ")
                professor = Professor(nome=nome, email=email)
                self.limpar_terminal()
                msg = es.cadastrar_professor(professor=professor)
                self.handle_response(msg)
            elif choice == '0':
                break
            
            else:
                print("Opção inválida. Tente novamente.")
                
    # implmenetar isso
    def handle_response(self,msg : Message):
        match msg.method:
            case "CadastrarAluno" | "CadastrarDisciplina" | "CadastrarProfessor":
                status = msg.arguments.get("status")
                mensagem = msg.arguments.get("message")
                print(f"Status: {status}")
                print(f"mensagem: {mensagem}")
            case "BuscarAlunoPorCodigo":
                status = msg.arguments.get("status")
                alunos = msg.arguments.get("alunos", [])
                print(f"Status: {status}")
                for aluno in alunos:
                    print(f"Nome: {aluno['nome']}, CPF: {aluno['cpf']}")
            case _:
                print("Erro")
                
    def limpar_terminal(self):
        os.system('cls' if os.name == 'nt' else 'clear')
