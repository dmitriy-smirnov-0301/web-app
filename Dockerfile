FROM golang:1.23.4-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ice-creams-app ./cmd/ice-creams-catalog/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates bash postgresql-client

COPY --from=builder /app/ice-creams-app .

ENV GIN_MODE=debug \
    APP_PORT=8080

EXPOSE 8080

CMD ["./ice-creams-app"]
