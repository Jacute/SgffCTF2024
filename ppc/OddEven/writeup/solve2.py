from pwn import *


io = remote('79.141.74.23', 11111)

while True:
    output = io.recvline()
    if b'SgffCTF' in output:
        print(output)
        break
    io.recvuntil(b': ') # skip line
    
    num = int(output[:-1])
    if num % 2 == 0:
        io.sendline(b'even')
    else:
        io.sendline(b'odd')
io.close()
