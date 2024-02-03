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
	Payload    string `json:"payload"`
	DelayNanos int64  `json:"delay_nanos"`
}

type PongMessage struct {
	Timestamp     string  `json:"timestamp"`
	Hostname      string  `json:"hostname"`
	Version       string  `json:"version"`
	RemoteIp      string  `json:"remote_ip"`
	RemotePort    int     `json:"remote_port"`
	RequestId     string  `json:"request_id"`
	RequestHost   string  `json:"request_host"`
	RequestMethod string  `json:"request_method"`
	RequestPath   string  `json:"request_path"`
	RequestQuery  string  `json:"request_query"`
	RequestTime   float64 `json:"request_time"`
	UserAgent     string  `json:"user_agent"`
	Payload       string  `json:"payload"`
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
		Payload:       ping.Payload,
	}
	return util.WriteJSON(w, pong)
}
