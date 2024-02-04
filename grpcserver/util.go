package grpcserver

import (
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
