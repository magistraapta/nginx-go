# Build stage
FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy only the binary from the builder stage
COPY --from=builder /app/main .
# Copy .env file and templates directory
COPY --from=builder /app/.env .
COPY --from=builder /app/templates ./templates

# Run as non-root user for better security
RUN adduser -D -g '' appuser
USER appuser

# Environment variables should be provided at runtime
# Example: docker run -e DB_HOST=localhost -e DB_PORT=5432 ...

EXPOSE 8080

CMD ["./main"]