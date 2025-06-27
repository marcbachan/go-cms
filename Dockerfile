# Use an official Go base image
FROM golang:1.24-alpine

# Enable Go modules and disable CGO
ENV CGO_ENABLED=0 GO111MODULE=on

# Set working directory in container
WORKDIR /app

# Copy go.mod and go.sum and download deps first (for cache efficiency)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app source code
COPY . .

# Create required folders (in case they don't exist)
RUN mkdir -p posts images

# Build the Go app
RUN go build -o cms-server .

# Expose the default port
EXPOSE 8080

# Run the server
CMD ["./cms-server"]
