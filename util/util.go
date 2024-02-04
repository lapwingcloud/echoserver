package util

import (
	"log/slog"
	"net"
	"os"
	"strings"
)

func NewLogger(logFormat string, commonArgs ...any) *slog.Logger {
	switch logFormat {
	case "json":
		return slog.New(slog.NewJSONHandler(os.Stderr, nil)).With(commonArgs...)
	default:
		return slog.Default().With(commonArgs...)
	}
}

func ParseNetAddr(addr net.Addr) (ip string, port int) {
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

func Version() string {
	version, _ := os.ReadFile("./version.txt")
	return strings.TrimSpace(string(version))
}
