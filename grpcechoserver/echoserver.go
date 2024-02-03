package grpcechoserver

import (
	"context"
	"strconv"
	"time"

	pb "github.com/lapwingcloud/echoserver/proto"
	"github.com/lapwingcloud/echoserver/util"
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
	remotePort, _ := strconv.Atoi(util.FirstValueFromMetadata(md, "remote-port"))
	startTime, err := time.Parse(time.RFC3339, util.FirstValueFromMetadata(md, "start-time"))
	var requestTime float64
	if err == nil {
		requestTime = time.Since(startTime).Seconds()
	}
	return &pb.PongMessage{
		Timestamp:     time.Now().Format(time.RFC3339),
		Hostname:      util.FirstValueFromMetadata(md, "hostname"),
		Version:       util.FirstValueFromMetadata(md, "version"),
		RemoteIp:      util.FirstValueFromMetadata(md, "remote-ip"),
		RemotePort:    int32(remotePort),
		RequestId:     util.FirstValueFromMetadata(md, "request-id"),
		Authority:     util.FirstValueFromMetadata(md, ":authority"),
		RequestMethod: util.FirstValueFromMetadata(md, "request-method"),
		RequestTime:   requestTime,
		UserAgent:     util.FirstValueFromMetadata(md, "user-agent"),
		Payload:       req.GetPayload(),
	}, nil
}
