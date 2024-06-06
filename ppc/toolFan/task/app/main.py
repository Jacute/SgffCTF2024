import json
import socket
from random import choice
import threading


def handle_client(client_socket, client_address):
        print(f"Connect from {client_address}")
        
        with open('data.json', 'r') as f:
            data = json.loads(f.read())
        
        c = 0
        while c != 500:
            elem = choice(data)
            question = elem["question"] + "\n"
            try:
                client_socket.send(question.encode('utf-8'))
                client_socket.send(b"Answer: ")
                user_answer = client_socket.recv(1024).decode('utf-8')
                if user_answer.lower()[:-1] == elem["answer"].lower():
                    client_socket.send(b"Correct!\n")
                    c += 1
                else:
                    client_socket.send(b"Incorrect! Right is " + elem["answer"].encode("utf-8") + b"\n")
            except BrokenPipeError:
                break
            except ConnectionResetError:
                break
        if c == 500:
            client_socket.send(b"SgffCTF{th3_b3st_t00l_f4n}\n")
        
        client_socket.close()
        print(f"Disconnected from {client_address}")


def main():
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    server_socket.bind(('0.0.0.0', 5553))
    server_socket.listen(5)
    
    while True:
        client_socket, client_address = server_socket.accept()
        client_handler = threading.Thread(target=handle_client, args=(client_socket , client_address, ))
        client_handler.start()


if __name__ == '__main__':
    main()