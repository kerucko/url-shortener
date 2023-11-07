FROM golang:1.21.0

WORKDIR /usr/src/shortener

COPY go.mod ./
RUN go mod download && go mod verify

COPY ./ ./
RUN go build -o shortener ./cmd/main.go

CMD ["./shortener"]