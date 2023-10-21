# Use the latest Go version (replace '1.20' with the appropriate version)
FROM golang:1.20-alpine as builder

WORKDIR /go/src/app
COPY . .

# Download dependencies and build the application
RUN go mod download
RUN go build -o vault-cluster-replication ./cmd/main.go

FROM golang:1.20-alpine

# Copy only the built binary from the builder stage
COPY --from=builder /go/src/app/vault-cluster-replication /app/vault-cluster-replication

# Create a group and user

RUN addgroup -S vcr && adduser -S vcr -G vcr
RUN chown -R vcr:vcr /app
USER vcr

# Set the working directory and entry point
WORKDIR /app
ENTRYPOINT ["/app/vault-cluster-replication"]

# Add labels for metadata
LABEL maintainer="Your Name <m.auso.p@gmail.com>"
LABEL version="1.0"
LABEL description="Vault Cluster Replication Application"
