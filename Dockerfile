# Use the official GoLang 1.20 image as the build environment
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o vault-cluster-replicator

# Use a minimal scratch image to reduce the container size
FROM scratch

# Copy the built binary from the builder stage
COPY --from=builder /app/vault-cluster-replicator /app

# Set the entry point for the container
ENTRYPOINT ["/vault-cluster-replicator"]
