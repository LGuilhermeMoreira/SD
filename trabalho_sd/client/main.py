from app.proxy.escola import EscolaService
from app.models import Aluno, Disciplina, Professor, Codigo

from enum import Enum

class MessageType(Enum):
    REQUEST = 0
    RESPONSE = 1

def main():
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
            es.buscar_aluno_por_codigo(codigo)
            # response = es.buscar_aluno_por_codigo(codigo)
            # if response:
            #   print(f"Resposta do servidor: {response}")
        
        elif choice == '2':
            nome = input("Nome do aluno: ")
            curso = input("Curso do aluno: ")
            cpf = input("CPF do aluno: ")
            disciplinas = []
            aluno = Aluno(nome=nome, curso=curso, cpf=cpf, disciplinas=disciplinas)
            es.cadastrar_aluno(aluno=aluno)
            # response = es.cadastrar_aluno(aluno=aluno)
            # if response:
            #   print(f"Resposta do servidor: {response}")

        elif choice == '3':
            nome = input("Nome da disciplina: ")
            codigo = input("Código da disciplina: ")
            professores = []
            disciplina = Disciplina(nome=nome, codigo=codigo, professores=professores)
            es.cadastrar_disciplina(disciplina=disciplina)
            # response = es.cadastrar_disciplina(disciplina=disciplina)
            # if response:
            #   print(f"Resposta do servidor: {response}")
        
        elif choice == '4':
            nome = input("Nome do professor: ")
            email = input("Email do professor: ")
            professor = Professor(nome=nome, email=email)
            es.cadastrar_professor(professor=professor)
            # response = es.cadastrar_professor(professor=professor)
            # if response:
            #    print(f"Resposta do servidor: {response}")
        
        elif choice == '0':
            break
        
        else:
            print("Opção inválida. Tente novamente.")

if __name__ == "__main__":
    main()