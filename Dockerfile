FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o vulscanner

FROM alpine:latest

RUN apk add --no-cache sqlite

WORKDIR /root/

COPY --from=builder /app/vulscanner .

COPY vulscanner.db .

EXPOSE 8080

# Run the application
CMD ["./vulscanner"]
