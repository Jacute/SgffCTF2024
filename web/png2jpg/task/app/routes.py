from . import app
from flask import render_template, request, send_file, jsonify, send_from_directory
from .utils import checkExstension, png2jpg

import os
import subprocess


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/convert/<string:file>')
def getImage(file):
    if file:
        return send_file(os.path.join('..', 'convert', file), mimetype='image/jpeg')

@app.route('/convert', methods=['POST'])
def convert():
    if 'image' not in request.files:
        return jsonify({"error": "File not uploaded"})
    else:
        file = request.files['image']
        filename = file.filename.replace(' ', '_')
        
        if filename == '':
            return jsonify({'error': "File not uploaded"}), 400
        else:
        
            if file and checkExstension(filename):
                filepath = os.path.join('uploads', filename)
                try:
                    file.save(filepath)
                except:
                    return jsonify({'error': "File can't be saved"}), 400
            else:
                return jsonify({'error': "File should be image"}), 400
    try:
        newFilePath = png2jpg(filename)
    except Exception as e:
        return jsonify({'error': f'Something went wrong. Error: {e}'})
    return jsonify({"file": newFilePath})
