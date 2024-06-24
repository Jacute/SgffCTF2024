import os


filespath = '../task/files'

files = os.listdir(filespath)
files.sort(key=lambda x: int(x.split('.')[0]))

buf = ''

for file in files:
    with open(filespath + '/' + file, 'r') as f:
        buf += f.read()
if 'SgffCTF' in buf:
    print('Flag here!')
    print(buf)
else:
    print('Flag not here!')
