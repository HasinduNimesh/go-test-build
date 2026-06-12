# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Download dependencies first (cached if go.mod/go.sum haven't changed)
COPY go.mod go.sum* ./
RUN go mod download

# Copy source and build a statically linked binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Stage 2: Minimal runtime
FROM scratch
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["/app/server"]