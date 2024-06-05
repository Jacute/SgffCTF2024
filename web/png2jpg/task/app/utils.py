import subprocess
import os

from . import app


def checkExstension(filename: str):
    return filename.lower().split('.')[-1] in app.config['ALLOWED_EXSTENSIONS']

def png2jpg(file: str):
    pngFilePath = os.path.join('uploads', file)
    jpgFilePath = os.path.join('convert', file.split(".")[0] + '.jpg')
    command = f'convert {pngFilePath} {jpgFilePath}'        
    subprocess.check_output(command, shell=True)
    
    return jpgFilePath
