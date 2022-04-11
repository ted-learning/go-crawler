FROM golang:1.17.8-alpine3.15

WORKDIR /app

COPY go.* ./

RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build

EXPOSE 7500

CMD ["/app/crawler-engine"]
