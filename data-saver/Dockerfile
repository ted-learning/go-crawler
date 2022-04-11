FROM golang:1.17.8-alpine3.15

WORKDIR /app

COPY go.* ./

RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build

EXPOSE 1234

CMD ["/app/data-saver"]
