package util

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
)

func FirstValueFromMetadata(md metadata.MD, key string) string {
	if md == nil {
		return ""
	}
	values := md.Get(key)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func Hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func NewLogger(logformat string) *slog.Logger {
	switch logformat {
	case "json":
		return slog.New(slog.NewJSONHandler(os.Stderr, nil))
	default:
		return slog.Default()
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

func WriteJSON(w http.ResponseWriter, data any) error {
	resp, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	resp = append(resp, byte('\n'))
	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(fmt.Sprintf(
		`{"timestamp":"%s","error":"%v"}`,
		time.Now().Format(time.RFC3339),
		err,
	)))
}

func Version() string {
	version, _ := os.ReadFile("./version.txt")
	return strings.TrimSpace(string(version))
}
