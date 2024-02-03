package grpcechoserver

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type unaryInterceptor struct {
	logger *slog.Logger
}

func (u *unaryInterceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()

	hostname, _ := os.Hostname()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || md == nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve metadata from incoming context")
	}
	p, _ := peer.FromContext(ctx)
	remoteIp, remotePort := parseNetAddr(p.Addr)
	serverIp, serverPort := parseNetAddr(p.LocalAddr)
	requestId := firstValueFromMetadata(md, "request-id")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	authority := firstValueFromMetadata(md, ":authority")
	userAgent := firstValueFromMetadata(md, "user-agent")

	md.Append("start-time", startTime.Format(time.RFC3339))
	md.Append("hostname", hostname)
	md.Append("remote-ip", remoteIp)
	md.Append("remote-port", fmt.Sprint(remotePort))
	md.Append("request-id", requestId)
	md.Append("request-method", info.FullMethod)

	ctx = metadata.NewIncomingContext(ctx, md)
	resp, err := handler(ctx, req)

	if err != nil {
		u.logger.Error(
			"grpc unary request error",
			"hostname", hostname,
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
			"hostname", hostname,
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
