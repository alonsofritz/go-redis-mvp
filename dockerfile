# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-redis-app

# Runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /go-redis-app /app/go-redis-app

EXPOSE 8080

CMD ["/app/go-redis-app"]