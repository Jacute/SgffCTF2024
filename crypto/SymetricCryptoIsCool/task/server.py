import socket
from threading import Thread
from random import shuffle, seed
import string
from seed import random_seed

def encode_message(message, table):
    encoded_message = []
    for char in message:
        for pair in table:
            if char == pair[0]:
                encoded_message.append(pair[1])
                break
            elif char == pair[1]:
                encoded_message.append(pair[0])
                break
        else:
            encoded_message.append(char)
    return ''.join(encoded_message)

# Initialize seed and create encoding table
seed(random_seed)
alphabet = [i for i in string.printable]
shuffled_alphabet = alphabet.copy()
shuffle(shuffled_alphabet)

encode_table = []
used = set()

for a in alphabet:
    if a in used:
        continue
    for b in shuffled_alphabet:
        if b not in used and a != b:
            encode_table.append([a, b])
            used.add(a)
            used.add(b)
            break

# Server details
HOST = '0.0.0.0'
PORT = 22222

menu = """\
Menu:
1. Encode a string
2. Exit
Choose an option: 
"""

def handle_client(conn, addr):
    print(f"Connected by {addr}")
    conn.sendall(menu.encode('utf-8'))
    while True:
        data = conn.recv(1024)
        if not data:
            break
        user_input = data.decode('utf-8').strip()
        
        if user_input == '1':
            conn.sendall(b'Enter a string to encode: ')
            data = conn.recv(1024)
            if not data:
                break
            message_to_encode = data.decode('utf-8').strip()
            encrypted_message = encode_message(message_to_encode, encode_table)
            response = f"Encoded message: {encrypted_message}\n\n{menu}"
            conn.sendall(response.encode('utf-8'))
        elif user_input == '2' or user_input.lower() == 'exit':
            conn.sendall(b'Goodbye!\n')
            break
        else:
            conn.sendall(b'Invalid option. Please choose 1 or 2.\n')
            conn.sendall(menu.encode('utf-8'))
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
