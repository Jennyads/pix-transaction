package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func ReadMetadata(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		return md.Get(key)[0]
	}
	return ""
}
