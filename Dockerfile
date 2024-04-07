# Use a smaller base image for the final stage
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Set environment variables
ENV PORT=8000
ENV GIN_MODE=release
ENV SHOPRITE=https://www.shoprite.co.za/search/all?q=
ENV CHECKERS=https://www.checkers.co.za/search/all?q=
ENV PNP=https://www.pnp.co.za/search/
ENV WOOLWORTHS=https://www.woolworths.co.za/cat?Dy=1&Ntt=
ENV MAKRO=https://www.makro.co.za
ENV GAME=https://www.game.co.za

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a smaller base image for the final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
