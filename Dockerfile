# Use the official Golang base image
FROM golang:1.20-alpine AS build

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code to the working directory
COPY . .

# Compile the application
RUN CGO_ENABLED=1 go build -o the-ethereum-fetcher .

# Use a lightweight base image for the final stage
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /app/the-ethereum-fetcher /app/the-ethereum-fetcher
COPY --from=build /app/config.yaml /app/config.yaml

# Expose the necessary ports
EXPOSE 2222
EXPOSE 5000
EXPOSE 9090

# Set environment variables
ENV THE_ETHEREUM_FETCHER_LIME_SERVER_LISTEN_ADDRESS=0.0.0.0:2222
ENV THE_ETHEREUM_FETCHER_HEALTHCHECK_SERVER_LISTEN_ADDRESS=0.0.0.0:5000
ENV THE_ETHEREUM_FETCHER_METRICS_SERVER_LISTEN_ADDRESS=0.0.0.0:9090
ENV THE_ETHEREUM_FETCHER_ETHEREUM_ADDRESS=https://goerli.infura.io/v3/0a3ab0b86d5e4835a3b94832195f4912
ENV THE_ETHEREUM_FETCHER_DB_POSTGRESQL_HOST=localhost
ENV THE_ETHEREUM_FETCHER_DB_POSTGRESQL_PORT=5432
ENV THE_ETHEREUM_FETCHER_DB_POSTGRESQL_SSL_MODE=disable
ENV THE_ETHEREUM_FETCHER_DB_POSTGRESQL_USERNAME=the-ethereum-fetcher
ENV THE_ETHEREUM_FETCHER_DB_POSTGRESQL_PASSWORD=Zgahbwm+qp3GNG4R
ENV THE_ETHEREUM_FETCHER_DB_POSTGRESQL_DATABASE=main
ENV THE_ETHEREUM_FETCHER_DB_POSTGRESQL_SCHEMA=public

# Run the application
CMD ["/app/the-ethereum-fetcher"]
