FROM golang:1.21 AS builder

COPY . /code
WORKDIR /code
RUN go build -o echoserver main.go

FROM debian:12

COPY --from=builder /code/echoserver /bin/echoserver
COPY version.txt version.txt
CMD ["echoserver"]
