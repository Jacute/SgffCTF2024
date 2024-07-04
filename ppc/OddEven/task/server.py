import socket
from threading import Thread
from random import randint


# Server details
HOST = '0.0.0.0'
PORT = 11111


def handle_client(conn, addr):
    print(f"Connected by {addr}")
    count = 0
    while True:
        number = randint(1, 100)
        # send number in UTF-8 format
        conn.sendall(str(number).encode('utf-8'))
        conn.sendall(b'\n')
        conn.sendall(b'Enter "odd" or "even": ')
        data = conn.recv(1024)
        if not data:
            break
        user_input = data.decode('utf-8').strip()
        
        if user_input == 'odd' and number % 2 != 0:
            count += 1
        elif user_input == 'even' and number % 2 == 0:
            count += 1
        else:
            conn.sendall(b'Wrong!\n')
            break
            
        if count == 100:
            conn.sendall(b'SgffCTF{N0t_B@D_For_fir$t_t1me}\n')
            break
    conn.close()
    print(f"Disconnected from {addr}")

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server_socket:
    server_socket.bind((HOST, PORT))
    server_socket.listen()
    print(f"Server listening on {HOST}:{PORT}")
    
    while True:
        conn, addr = server_socket.accept()
        client_thread = Thread(target=handle_client, args=(conn, addr))
        client_thread.start()
