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
COPY --from=builder /app/.env .

# Run as non-root user for better security
RUN adduser -D -g '' appuser
USER appuser


EXPOSE 8080

CMD ["./main"]