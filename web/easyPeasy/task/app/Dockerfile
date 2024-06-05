FROM golang:1.22.3


WORKDIR /app

COPY . .

RUN go get
RUN go build main.go

CMD ["./main"]
