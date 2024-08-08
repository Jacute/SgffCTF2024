FROM golang:1.22.5


WORKDIR /app

COPY . .

RUN go get
RUN go build main.go

CMD ["./main"]
