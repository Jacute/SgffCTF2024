FROM python:3.10-alpine

WORKDIR /app

COPY templates templates
COPY main.py .

RUN pip install --no-cache-dir Flask==3.0.3

CMD ["python3", "main.py"]
