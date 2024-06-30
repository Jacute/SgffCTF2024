import socket

# Server details
HOST = '127.0.0.1'  # Или IP-адрес сервера
PORT = 11111

def main():
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
        client_socket.connect((HOST, PORT))
        print("Connected to the server")

        while True:
            # Получаем данные от сервера
            data = client_socket.recv(1024).decode('utf-8')
            if not data:
                break

            # Разделяем строки по новой строке
            lines = data.split('\n')
            for line in lines:
                line = line.strip()
                if line.isdigit():
                    number = int(line)
                    if number % 2 == 0:
                        client_socket.sendall(b'even\n')
                    else:
                        client_socket.sendall(b'odd\n')
                elif line == 'Enter "odd" or "even":':
                    continue
                else:
                    print(line)

if __name__ == "__main__":
    main()
