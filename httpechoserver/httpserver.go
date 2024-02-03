package httpechoserver

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type httpServer struct {
	logger     *slog.Logger
	statusCode int
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	s.statusCode = 0

	hostname, _ := os.Hostname()
	remoteIp, remotePortStr, _ := net.SplitHostPort(r.RemoteAddr)
	remotePort, _ := strconv.Atoi(remotePortStr)
	requestId := r.Header.Get("X-Request-Id")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	userAgent := r.Header.Get("User-Agent")

	requestContext := &RequestContext{
		StartTime:  startTime,
		Hostname:   hostname,
		RemoteIp:   remoteIp,
		RemotePort: remotePort,
		RequestId:  requestId,
		UserAgent:  userAgent,
	}
	ctx := context.WithValue(r.Context(), requestContextKey, requestContext)
	err := ping(w, r.WithContext(ctx))
	if err != nil {
		writeError(w, err)
		s.logger.Error(
			"http request error",
			"hostname", hostname,
			"remote_ip", remoteIp,
			"remote_port", remotePort,
			"request_id", requestId,
			"request_method", r.Method,
			"request_host", r.Host,
			"request_path", r.URL.Path,
			"request_query", r.URL.RawQuery,
			"request_time", time.Since(startTime).Seconds(),
			"user_agent", userAgent,
			"status", http.StatusInternalServerError,
			"error", err,
		)
	} else {
		s.logger.Info(
			"http request ok",
			"hostname", hostname,
			"remote_ip", remoteIp,
			"remote_port", remotePort,
			"request_id", requestId,
			"request_method", r.Method,
			"request_host", r.Host,
			"request_path", r.URL.Path,
			"request_query", r.URL.RawQuery,
			"request_time", time.Since(startTime).Seconds(),
			"user_agent", userAgent,
			"status", http.StatusOK,
		)
	}
}
