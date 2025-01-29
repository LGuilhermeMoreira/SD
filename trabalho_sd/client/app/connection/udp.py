import socket
from app.models import Message


def enviar_mensagem(host: str = 'localhost', port: int = 4567, message: Message = None):
    addr = (host, port)
    try:
        with socket.socket(socket.AF_INET, socket.SOCK_DGRAM) as sock:
            message_json = message.to_json()
            sock.sendto(message_json.encode('utf-8'), addr)
            sock.settimeout(5)  
            data, _ = sock.recvfrom(1024)
            response = message.from_json(data)
            return response

    except socket.timeout:
        print("Erro: Timeout ao receber resposta do servidor.")
        return None
    except Exception as e:
        print(f"Erro ao enviar/receber: {e}")
        return None