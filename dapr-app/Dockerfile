
FROM golang:1.17-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/hellodapr-go

COPY main.go .
COPY go.mod .
COPY go.sum .

# Build the Go app
RUN go build -o ./out/hellodapr-go .

# Start fresh from a smaller image
FROM alpine:3.14
RUN apk add ca-certificates

COPY --from=build_base /tmp/hellodapr-go/out/hellodapr-go /app/hellodapr-go

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/hellodapr-go"]