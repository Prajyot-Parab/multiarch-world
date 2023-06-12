# Dockerfile

# Use the official golang base image
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the source code into the container
COPY main.go main.go

# Build the Go application
RUN CGO_ENABLED=0 go build -o multiarch-world .

# Set the command to run the executable
CMD ["/app/multiarch-world"]
