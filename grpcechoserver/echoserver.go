package grpcechoserver

import (
	"context"
	"strconv"
	"time"

	pb "github.com/lapwingcloud/echoserver/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type echoServer struct {
	pb.UnimplementedEchoServer
}

func (s *echoServer) Ping(ctx context.Context, req *pb.PingMessage) (*pb.PongMessage, error) {
	if req.GetDelayNanos() > 0 {
		time.Sleep(time.Duration(req.GetDelayNanos()) * time.Nanosecond)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || md == nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve metadata from incoming context")
	}
	remotePort, _ := strconv.Atoi(firstValueFromMetadata(md, "remote-port"))
	startTime, err := time.Parse(time.RFC3339, firstValueFromMetadata(md, "start-time"))
	var requestTime float64
	if err == nil {
		requestTime = time.Since(startTime).Seconds()
	}
	return &pb.PongMessage{
		Timestamp:     time.Now().Format(time.RFC3339),
		Hostname:      firstValueFromMetadata(md, "hostname"),
		RemoteIp:      firstValueFromMetadata(md, "remote-ip"),
		RemotePort:    int32(remotePort),
		RequestId:     firstValueFromMetadata(md, "request-id"),
		Authority:     firstValueFromMetadata(md, ":authority"),
		RequestMethod: firstValueFromMetadata(md, "request-method"),
		RequestTime:   requestTime,
		UserAgent:     firstValueFromMetadata(md, "user-agent"),
		Payload:       req.GetPayload(),
	}, nil
}
