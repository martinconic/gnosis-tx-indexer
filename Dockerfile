FROM golang:1.21.5 as builder

WORKDIR /indexer

COPY . . 

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/.

CMD ["/indexer/main/main"]