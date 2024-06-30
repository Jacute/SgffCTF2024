from Crypto.Util.number import bytes_to_long, getPrime


flag = bytes_to_long(b'SgffCTF(^_^ REDACTED ^_^)')
p = getPrime(2048)
q = getPrime(2048)
n = p * q
e = 3
enc = pow(flag, e, n)
print(enc)
#настоящий зашифрованный флаг смотри в файле task.txt
