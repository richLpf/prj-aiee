# Use the official Golang image as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
# COPY go.mod go.sum ./

# Download and install the Go dependencies
# RUN go mod download

# Copy the source code into the container
COPY ./app .

# Build the Go application
# RUN go build 

# Set the entry point for the container
ENTRYPOINT ["./app"]
