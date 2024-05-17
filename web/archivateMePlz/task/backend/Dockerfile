FROM golang:1.22.2


WORKDIR /app

COPY . .

RUN go get
RUN go build main.go cleaner.go

CMD ["./main"]
