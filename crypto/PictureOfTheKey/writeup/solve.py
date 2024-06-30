from Crypto.Util.number import long_to_bytes


with open("crypto/PictureOfTheKey/writeup/flag.enc", "rb") as f:
    key = f.read()

with open("crypto/PictureOfTheKey/writeup/key.jpg", "rb") as f:
    flag = f.read()

encrypted = b''
for key_byte, flag_byte in zip(key, flag):
    encrypted += long_to_bytes(key_byte ^ flag_byte)

with open("crypto/PictureOfTheKey/writeup/flag.jpg", "wb") as f:
    f.write(encrypted)
