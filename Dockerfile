FROM golang:1.21 AS builder

COPY . /code
WORKDIR /code
RUN CGO_ENABLED=0 GOOS=linux go build -o echoserver main.go

FROM digitalocean/doks-debug

COPY --from=builder /code/echoserver echoserver
COPY version.txt version.txt
CMD ./echoserver
