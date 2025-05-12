# Use the official Go image to build our executable.
FROM golang:1.21-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /functions

# Copy go mod and go.sum to download all dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container to compile the app
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal Alpine image to run the application
FROM alpine:latest

# Copy binary from builder stage
COPY --from=builder /functions/main /functions/main

# Set the Current Working Directory inside the container
WORKDIR /functions

# Expose the port that the application listens on
EXPOSE 8080

# Command to run the executable
CMD ["/functions/main"]