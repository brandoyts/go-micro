FROM golang:1.23.10 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-gateway ./api-gateway/cmd/server

FROM alpine:latest

RUN adduser -D -g '' brando

COPY --from=builder /api-gateway /api-gateway

USER brando

ENTRYPOINT ["/api-gateway"]