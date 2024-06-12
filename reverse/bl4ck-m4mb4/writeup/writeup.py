with open('encrypted.txt', 'rb') as f:
    data = f.read()
    
for i in range(len(data)):
    print(chr(data[i] - i), end='')
print()
