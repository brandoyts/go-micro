FROM golang:1.23.10 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth ./auth/cmd/server

FROM alpine:latest

RUN adduser -D -g '' brando

COPY --from=builder /auth /auth

USER brando

ENTRYPOINT ["/auth"]