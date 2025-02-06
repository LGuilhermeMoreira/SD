import socket

class UDPCliente:
    
    def __init__(self):
        self.addr = ('localhost', 4567)
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        self.socket.settimeout(1) 
        
        ok = self._ping_no_servidor()
        if not ok:
            raise ConnectionError("Não foi possível se conectar ao servidor")
        
    def enviar_mensagem(self,message_json: str = ""):    
        self.socket.sendto(message_json.encode('utf-8'),self.addr)
        
    def receber_mensagem(self):
        data, _ = self.socket.recvfrom(1024)
        return data

    def _ping_no_servidor(self):
        try:
            self.enviar_mensagem()
            self.receber_mensagem()
        except socket.timeout:
            return False
        except Exception as e:
            print(f"Erro ao tentar pingar o servidor: {e}")
            return False