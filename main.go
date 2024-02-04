package main

import (
	"flag"
	"sync"

	"github.com/lapwingcloud/echoserver/grpcserver"
	"github.com/lapwingcloud/echoserver/httpserver"
)

var (
	grpcBind  = flag.String("grpc-bind", ":9090", "The grpc server listen address")
	httpBind  = flag.String("http-bind", ":8080", "The http server listen address")
	logFormat = flag.String("log-format", "json", "The log format (text, json)")
)

func main() {
	flag.Parse()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		grpcserver.Start(grpcserver.StartOption{
			Bind:      *grpcBind,
			LogFormat: *logFormat,
		})
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		httpserver.Start(httpserver.StartOption{
			Bind:      *httpBind,
			LogFormat: *logFormat,
		})
		wg.Done()
	}()

	wg.Wait()
}
