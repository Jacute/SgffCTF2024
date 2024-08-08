FROM golang:1.22.5


RUN adduser -S -D -H -h /app app
RUN echo "app:x:1000:1000::/app:/sbin/nologin" >> /etc/passwd

WORKDIR /app
COPY . .

RUN chown app *

RUN go get
RUN go build main.go

USER app

CMD ["./main"]
