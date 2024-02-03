package main

import (
	"flag"
	"sync"

	"github.com/lapwingcloud/echoserver/grpcechoserver"
	"github.com/lapwingcloud/echoserver/httpechoserver"
)

var (
	grpcBind  = flag.String("grpc-bind", "127.0.0.1:9090", "The grpc server listen address")
	httpBind  = flag.String("http-bind", "127.0.0.1:8080", "The http server listen address")
	logFormat = flag.String("log-format", "json", "The log format (text, json)")
)

func main() {
	flag.Parse()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		grpcechoserver.Start(grpcechoserver.StartOption{
			Bind:      *grpcBind,
			LogFormat: *logFormat,
		})
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		httpechoserver.Start(httpechoserver.StartOption{
			Bind:      *httpBind,
			LogFormat: *logFormat,
		})
		wg.Done()
	}()

	wg.Wait()
}
