
# Use the official Go image as a base image
FROM golang:1.23 as builder

# Set environment variables
ENV GO111MODULE=on     CGO_ENABLED=0     GOOS=linux     GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o oriva cmd/main.go

# Use a minimal image for the final build
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory inside the final container
WORKDIR /root/

# Copy the binary from the builder image
COPY --from=builder /app/oriva .

# Copy the .env file if needed
COPY .env .

# Expose the port that the gRPC server listens on
EXPOSE 50027

# Command to run the application
CMD ["./oriva"]
