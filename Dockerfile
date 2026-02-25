# -------- Stage 1: Build Stage --------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy remaining source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server


# -------- Stage 2: Runtime Stage --------
FROM alpine:latest

WORKDIR /root/

# Copy compiled binary from builder stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]