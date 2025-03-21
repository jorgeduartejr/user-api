# Use the base image for Golang
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum (if they exist)
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o user-api .

# Use a minimal image for execution
FROM alpine:latest

# Install certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from the previous stage
COPY --from=builder /app/user-api .

# OpenShift runs containers with a random UID, so we need to fix permissions
RUN chown -R 1001:0 /root && \
    chmod -R g=u /root

# Use non-root user
USER 1001

# Expose the port used by the application
EXPOSE 3000

# Define default environment variables
ENV USE_MONGODB="false"
ENV MONGODB_URI="mongodb://localhost:27017"
ENV DATABASE_NAME="userdb"

# Command to start the application
CMD ["./user-api"]