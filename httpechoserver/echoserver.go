package httpechoserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lapwingcloud/echoserver/util"
)

type PingMessage struct {
	DelayNanos int64  `json:"delayNanos,omitempty"`
	Payload    string `json:"payload,omitempty"`
}

type PongMessage struct {
	Timestamp     string  `json:"timestamp"`
	Hostname      string  `json:"hostname"`
	Version       string  `json:"version"`
	RemoteIp      string  `json:"remoteIp"`
	RemotePort    int     `json:"remotePort"`
	RequestId     string  `json:"requestId"`
	RequestHost   string  `json:"requestHost"`
	RequestMethod string  `json:"requestMethod"`
	RequestPath   string  `json:"requestPath"`
	RequestQuery  string  `json:"requestQuery"`
	RequestTime   float64 `json:"requestTime"`
	UserAgent     string  `json:"userAgent"`
	PingMessage
}

func ping(w http.ResponseWriter, r *http.Request) error {
	requestContext, ok := r.Context().Value(requestContextKey).(*RequestContext)
	if !ok || requestContext == nil {
		return errors.New("failed to retrieve request context")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %v", err)
	}

	var ping PingMessage
	if len(body) != 0 {
		err := json.Unmarshal(body, &ping)
		if err != nil {
			return fmt.Errorf("failed to parse request body: %v", err)
		}
	}

	if ping.DelayNanos > 0 {
		time.Sleep(time.Duration(ping.DelayNanos) * time.Nanosecond)
	}

	pong := PongMessage{
		Timestamp:     time.Now().Format(time.RFC3339),
		Hostname:      requestContext.Hostname,
		Version:       requestContext.Version,
		RemoteIp:      requestContext.RemoteIp,
		RemotePort:    requestContext.RemotePort,
		RequestId:     requestContext.RequestId,
		RequestMethod: r.Method,
		RequestHost:   r.Host,
		RequestPath:   r.URL.Path,
		RequestQuery:  r.URL.RawQuery,
		RequestTime:   time.Since(requestContext.StartTime).Seconds(),
		UserAgent:     requestContext.UserAgent,
		PingMessage: PingMessage{
			Payload:    ping.Payload,
			DelayNanos: ping.DelayNanos,
		},
	}
	return util.WriteJSON(w, pong)
}
