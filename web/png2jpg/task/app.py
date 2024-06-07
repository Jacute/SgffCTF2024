import os
from app import app


if __name__ == '__main__':
    try:
        os.mkdir('uploads')
        os.mkdir('convert')
    except Exception:
        pass
    
    app.run(debug=False, host="0.0.0.0", port=8084)
