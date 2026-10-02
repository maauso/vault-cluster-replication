# Use the latest Go version (replace '1.24' with the appropriate version)
FROM golang:1.24-alpine as builder

WORKDIR /go/src/app
COPY . .

# Download dependencies and build the application
RUN go mod download \
    && go build -o vault-cluster-replication ./cmd/main.go

FROM golang:1.24-alpine

# Copy only the built binary from the builder stage
COPY --from=builder /go/src/app/vault-cluster-replication /app/vault-cluster-replication

# Create a group and user

RUN addgroup -S vcr && adduser -S vcr -G vcr \
    && chown -R vcr:vcr /app
USER vcr

# Set the working directory and entry point
WORKDIR /app
ENTRYPOINT ["/app/vault-cluster-replication"]

# Add labels for metadata
LABEL maintainer="Miguel Ausó"
LABEL version="1.1.0"
LABEL description="Vault Cluster Replication Application"
