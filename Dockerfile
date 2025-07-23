FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server && chmod +x server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

CMD ["./server"]