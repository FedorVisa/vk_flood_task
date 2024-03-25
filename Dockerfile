FROM golang:latest

LABEL authors="kotozavr"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

ENV REDIS_ADDR=redis:6379

CMD ["./main"]
