FROM golang:1.20-buster as builder

# Create and change to the app directory.
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

# Build the executable file.
RUN go build -v -o tweets

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y postgresql-client

WORKDIR /bin

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/tweets /bin/tweets

# Run the web service on container startup.
ENTRYPOINT [ "/bin/tweets" ]
