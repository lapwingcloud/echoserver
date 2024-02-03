package grpcechoserver

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/lapwingcloud/echoserver/proto"
	"github.com/lapwingcloud/echoserver/util"
	"google.golang.org/grpc"
)

type StartOption struct {
	Bind      string
	LogFormat string
}

func Start(option StartOption) {
	logger := util.NewLogger(option.LogFormat)

	lis, err := net.Listen("tcp", option.Bind)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to listen: %v", err))
		os.Exit(1)
	}

	ui := unaryInterceptor{
		logger:   logger,
		hostname: util.Hostname(),
		version:  util.Version(),
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(ui.Intercept))
	pb.RegisterEchoServer(s, &echoServer{})

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
