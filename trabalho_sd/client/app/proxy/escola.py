from app.models import Aluno, Codigo, Disciplina, Professor, Message
from app.connection import enviar_mensagem

class EscolaService:
    
    def _tentar_enviar_mensagem(self, msg: Message) -> Message | None:
        msg_json = msg.to_json()
        tentativas = 3
        for _ in range(tentativas):
            data = enviar_mensagem(message_json=msg_json)
            if data is not None:
                return Message.from_json(data)
        return None
    
    def cadastrar_aluno(self, aluno: Aluno) -> Message | None:
        msg = Message("Escola", "CadastrarAluno", aluno.__dict__)
        return self._tentar_enviar_mensagem(msg)
    
    def cadastrar_disciplina(self, disciplina: Disciplina) -> Message | None:
        msg = Message("Escola", "CadastrarDisciplina", disciplina.__dict__)
        return self._tentar_enviar_mensagem(msg)
    
    def cadastrar_professor(self, professor: Professor) -> Message | None:
        msg = Message("Escola", "CadastrarProfessor", professor.__dict__)
        return self._tentar_enviar_mensagem(msg)
    
    def buscar_aluno_por_codigo(self, codigo: Codigo) -> Message | None:
        msg = Message("Escola", "BuscarAlunoPorCodigo", codigo.__dict__)
        return self._tentar_enviar_mensagem(msg)