FROM golang:1.21-alpine AS builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
RUN apk add --no-cache git

# Clone the repository from GitHub
RUN git clone https://github.com/ErickAguilar95/golang-simple-api.git .

# Initialize Go module
RUN go mod init golang-simple-api

# Copy the source code
COPY . .

# Build the application
RUN go build -o /build/app .

# Stage 2: Final image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /build/app .

# Expose the port the application runs on
EXPOSE 8000

# Command to run the application
CMD ["./app"]
