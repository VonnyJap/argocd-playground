# Start from the official Golang base image
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o temperature-server .

# Use distroless as the final base image
FROM gcr.io/distroless/static-debian11

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/temperature-server .

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./temperature-server"]