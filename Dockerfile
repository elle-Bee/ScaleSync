# Use an official Golang image as the base
FROM golang:1.23-alpine

WORKDIR /app

# Copy the Go application files into the container
COPY . .

# Download Go module dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the Go app port (8080)
EXPOSE 8080

# Run the Go application
CMD ["./main"]
