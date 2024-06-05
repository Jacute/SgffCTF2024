from flask import Flask


app = Flask(__name__)


app.config['MAX_CONTENT_LENGTH'] = 20 * 1024 * 1024

app.config['ALLOWED_EXSTENSIONS'] = {'png'}

from . import routes