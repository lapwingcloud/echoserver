package grpcechoserver

import (
	"fmt"
	"net"
	"os"

	pb "github.com/lapwingcloud/echoserver/proto"
	"google.golang.org/grpc"
)

type StartOption struct {
	Bind      string
	LogFormat string
}

func Start(option StartOption) {
	logger := newLogger(option.LogFormat)

	lis, err := net.Listen("tcp", option.Bind)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to listen: %v", err))
		os.Exit(1)
	}

	ui := unaryInterceptor{
		logger: logger,
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(ui.Intercept))
	pb.RegisterEchoServer(s, &echoServer{})
	logger.Info(fmt.Sprintf("grpc server listening at %v", lis.Addr()))

	err = s.Serve(lis)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(1)
	}
}
