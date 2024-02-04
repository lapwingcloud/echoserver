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
	version  string
	hostname string
}

func (u *unaryInterceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || md == nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve metadata from incoming context")
	}
	p, _ := peer.FromContext(ctx)
	remoteIp, remotePort := util.ParseNetAddr(p.Addr)
	requestId := util.FirstValueFromMetadata(md, "request-id")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	authority := util.FirstValueFromMetadata(md, ":authority")
	userAgent := util.FirstValueFromMetadata(md, "user-agent")

	md.Append("start-time", startTime.Format(time.RFC3339Nano))
	md.Append("version", u.version)
	md.Append("hostname", u.hostname)
	md.Append("remote-ip", remoteIp)
	md.Append("remote-port", fmt.Sprint(remotePort))
	md.Append("request-id", requestId)
	md.Append("request-method", info.FullMethod)

	ctx = metadata.NewIncomingContext(ctx, md)
	resp, err := handler(ctx, req)

	u.logger.Info(
		"grpc unary request finished",
		slog.String("remoteIp", remoteIp),
		slog.Int("remotePort", remotePort),
		slog.String("requestId", requestId),
		slog.String("authority", authority),
		slog.String("requestMethod", info.FullMethod),
		slog.Float64("requestTime", time.Since(startTime).Seconds()),
		slog.String("userAgent", userAgent),
		slog.Int("status", int(status.Code(err))),
	)
	if err != nil {
		u.logger.Error(
			"grpc unary request error",
			slog.String("requestId", requestId),
			slog.String("error", err.Error()),
		)
	}
	return resp, err
}
