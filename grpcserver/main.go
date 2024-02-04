package grpcserver

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/lapwingcloud/echoserver/proto"
	"github.com/lapwingcloud/echoserver/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type StartOption struct {
	Bind      string
	LogFormat string
}

func Start(option StartOption) {
	version := util.Version()
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("failed to retrieve hostname: %v", err)
	}
	logger := util.NewLogger(
		option.LogFormat,
		slog.String("version", version),
		slog.String("hostname", hostname),
	)

	lis, err := net.Listen("tcp", option.Bind)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to listen: %v", err))
		os.Exit(1)
	}

	ui := unaryInterceptor{
		logger:   logger,
		version:  version,
		hostname: hostname,
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(ui.Intercept))
	pb.RegisterEchoServer(s, &echoServer{})
	reflection.Register(s)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-ch
		logger.Info(fmt.Sprintf("got signal %v, shutting down grpc server", sig))
		s.GracefulStop()
	}()

	logger.Info(fmt.Sprintf("grpc server listening at %v", lis.Addr()))
	err = s.Serve(lis)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(1)
	}
	logger.Info("grpc server has shut down")
}
