FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o turfly ./cmd/server

# Expose the port that the app will run on
EXPOSE 8080

# Command to run the application
CMD ["./turfly"]
