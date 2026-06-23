# 1. Build stage
# ------------------------
# 1. Build stage
# ------------------------
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Better caching (more stable builds)
COPY go.mod ./
RUN go mod download

COPY . .

# Production-safe build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o kubetask ./cmd/server


# ------------------------
# 2. Runtime stage
# ------------------------
FROM alpine:latest

# Security: CA certificates (needed for HTTPS calls)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary only (minimal image)
COPY --from=builder /app/kubetask .

EXPOSE 8080

# Non-root execution (safer)
RUN adduser -D appuser
USER appuser

CMD ["./kubetask"]