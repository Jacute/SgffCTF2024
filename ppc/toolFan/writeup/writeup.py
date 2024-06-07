from pwn import *
import re


dct = {}


ip = '79.141.65.118'
port = 5553
io = remote(ip, port)

# get dictionary
for i in range(200):
    question = io.recvline()
    io.sendline(b"1")
    answer = re.search(r'is\s+(.+)', io.recvline().decode()).group(1)
    dct[question] = answer.encode()

# get flag
while True:
    question = io.recvline()
    print(question)
    io.recv()
    io.sendline(dct[question])
    status = io.recvline()
    print(status)
