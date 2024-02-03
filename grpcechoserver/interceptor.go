package grpcechoserver

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/lapwingcloud/echoserver/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type unaryInterceptor struct {
	logger   *slog.Logger
	hostname string
	version  string
}

func (u *unaryInterceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || md == nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve metadata from incoming context")
	}
	p, _ := peer.FromContext(ctx)
	remoteIp, remotePort := util.ParseNetAddr(p.Addr)
	serverIp, serverPort := util.ParseNetAddr(p.LocalAddr)
	requestId := util.FirstValueFromMetadata(md, "request-id")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	authority := util.FirstValueFromMetadata(md, ":authority")
	userAgent := util.FirstValueFromMetadata(md, "user-agent")

	md.Append("start-time", startTime.Format(time.RFC3339))
	md.Append("hostname", u.hostname)
	md.Append("version", u.version)
	md.Append("remote-ip", remoteIp)
	md.Append("remote-port", fmt.Sprint(remotePort))
	md.Append("request-id", requestId)
	md.Append("request-method", info.FullMethod)

	ctx = metadata.NewIncomingContext(ctx, md)
	resp, err := handler(ctx, req)

	if err != nil {
		u.logger.Error(
			"grpc unary request error",
			"hostname", u.hostname,
			"version", u.version,
			"server_ip", serverIp,
			"server_port", serverPort,
			"remote_ip", remoteIp,
			"remote_port", remotePort,
			"request_id", requestId,
			"authority", authority,
			"request_method", info.FullMethod,
			"request_time", time.Since(startTime).Seconds(),
			"user_agent", userAgent,
			"status", status.Code(err),
			"error", err,
		)
	} else {
		u.logger.Info(
			"grpc unary request ok",
			"hostname", u.hostname,
			"version", u.version,
			"server_ip", serverIp,
			"server_port", serverPort,
			"cilent_ip", remoteIp,
			"client_port", remotePort,
			"request_id", requestId,
			"authority", authority,
			"request_method", info.FullMethod,
			"request_time", time.Since(startTime).Seconds(),
			"user_agent", userAgent,
			"status", status.Code(err),
		)
	}

	return resp, err
}
