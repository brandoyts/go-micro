FROM golang:1.23.10 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /user ./user

FROM alpine:latest

RUN adduser -D -g '' brando

COPY --from=builder /user /user

USER brando

ENTRYPOINT ["/user", "--migrate", "--seed"]