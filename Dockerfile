# Use the official Golang image with the required version
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Expose the port your app runs on
# EXPOSE 7000

# Command to run the application
CMD ["go", "run", "main.go"]