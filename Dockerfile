# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git for fetching dependencies
RUN apk add --no-cache git

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o math-mcp-server .

# Runtime stage
FROM alpine:3.23

WORKDIR /app

# Add ca-certificates for HTTPS (if needed)
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/math-mcp-server .

# Environment variables for configuration
ENV MATH_CATEGORIES=all
ENV TRANSPORT=http

# Expose HTTP port
EXPOSE 8080

# Run as non-root user for security
RUN adduser -D -g '' appuser
USER appuser

# Default command
ENTRYPOINT ["./math-mcp-server"]
