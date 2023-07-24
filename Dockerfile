# Start from a Debian based image with Go installed
FROM golang:1.20

# Install make
RUN apt-get update && apt-get install -y make

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed 
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container 
COPY . .

# Build the Go app
RUN make init

# Expose port 8080 to the Docker host, so we can access it from the outside.
EXPOSE 8080

# Command to run the executable
CMD ["make", "launch"]