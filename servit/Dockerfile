# Start from the official Go image (uses latest stable Go)
FROM golang

# Set working directory inside container
WORKDIR /servit


COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary inside the container
RUN go build -o servit

# Command to run the executable
CMD ["./myapp"]
