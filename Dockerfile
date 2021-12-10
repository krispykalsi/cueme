FROM golang:latest

RUN mkdir -p /go/src/cueme

WORKDIR /go/src/cueme

COPY go.mod /go/src/cueme
COPY go.sum /go/src/cueme
RUN go mod download

COPY . /go/src/cueme
RUN go build -o bin/server ./cmd/cueme-server

CMD ["./bin/server"]

EXPOSE 8080