# Build stage
## Pull golang image from the hub
FROM golang:alpine AS builder

## Set up ENV vars
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

## Choose work directory
WORKDIR /syns-users-ms

## Copy local project to docker container
COPY . .

## Run build command
RUN go build -o syns-users .

# Run stage
## alpine:latest image is a light linux image
FROM alpine:latest AS runner

## Choose work directory
WORKDIR /syns-users-ms

## Copy the executable binary file and .env file from the last stage to the new stage
COPY --from=builder /syns-users-ms/syns-users .

# Execute the build
CMD ["./syns-users"]