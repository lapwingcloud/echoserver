package httpserver

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type httpServer struct {
	logger *slog.Logger

	version  string
	hostname string
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	remoteIp, remotePortStr, _ := net.SplitHostPort(r.RemoteAddr)
	remotePort, _ := strconv.Atoi(remotePortStr)
	requestId := r.Header.Get("X-Request-Id")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	userAgent := r.Header.Get("User-Agent")

	requestContext := &RequestContext{
		StartTime:  startTime,
		Version:    s.version,
		Hostname:   s.hostname,
		RemoteIp:   remoteIp,
		RemotePort: remotePort,
		RequestId:  requestId,
		UserAgent:  userAgent,
	}
	ctx := context.WithValue(r.Context(), requestContextKey, requestContext)
	err := ping(w, r.WithContext(ctx))
	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusInternalServerError
	}
	s.logger.Info(
		"http request finished",
		slog.String("remoteIp", remoteIp),
		slog.Int("remotePort", remotePort),
		slog.String("requestId", requestId),
		slog.String("requestMethod", r.Method),
		slog.String("requestHost", r.Host),
		slog.String("requestPath", r.URL.Path),
		slog.String("requestQuery", r.URL.RawQuery),
		slog.Float64("requestTime", time.Since(startTime).Seconds()),
		slog.String("userAgent", userAgent),
		slog.Int("status", statusCode),
	)
	if err != nil {
		s.logger.Error(
			"http request error",
			slog.String("requestId", requestId),
			slog.String("error", err.Error()),
		)
		writeError(w, err)
	}
}
