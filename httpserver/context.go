package httpserver

import "time"

type ContextKey int

const requestContextKey ContextKey = 0

type RequestContext struct {
	StartTime  time.Time
	Hostname   string
	Version    string
	RemoteIp   string
	RemotePort int
	RequestId  string
	UserAgent  string
}
