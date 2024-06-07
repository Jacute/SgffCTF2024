FROM python:3.10

WORKDIR /app

RUN adduser user

COPY ./app ./app
COPY requirements.txt .

RUN mkdir uploads convert

RUN pip install --no-cache-dir -r requirements.txt

RUN chown user /app
RUN chown user /app/*

USER user

CMD ["gunicorn", "-w", "4", "-b", "0.0.0.0:8084", "app:app"]
