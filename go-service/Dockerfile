# ============================
# Stage 1 — Build binary
# ============================
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# ============================
# Stage 2 — Minimal runtime
# ============================
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# No shell, no dev tools, no surprises
# Expose the port your service uses
