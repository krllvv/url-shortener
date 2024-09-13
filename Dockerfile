FROM golang:1.22.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url-shortener ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/url-shortener /app/
COPY .env.example .env

CMD ["./url-shortener", "-d"]