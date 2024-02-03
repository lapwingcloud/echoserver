package grpcechoserver

import (
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc/metadata"
)

func firstValueFromMetadata(md metadata.MD, key string) string {
	if md == nil {
		return ""
	}
	values := md.Get(key)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func newLogger(logformat string) *slog.Logger {
	switch logformat {
	case "json":
		return slog.New(slog.NewJSONHandler(os.Stderr, nil))
	default:
		return slog.Default()
	}
}

func parseNetAddr(addr net.Addr) (ip string, port int) {
	switch a := addr.(type) {
	case *net.TCPAddr:
		ip = a.IP.String()
		port = a.Port
	case *net.UDPAddr:
		ip = a.IP.String()
		port = a.Port
	}
	return
}
