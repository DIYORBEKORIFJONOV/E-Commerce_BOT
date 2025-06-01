FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY .env .env
COPY auth.conf auth.conf
COPY auth.csv auth.csv

COPY --from=builder /app/main .

CMD ["./main"]
