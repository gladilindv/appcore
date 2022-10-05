package mw

import (
	"context"

	"github.com/gladilindv/appcore/recovery"
	"google.golang.org/grpc"
)

// RecoverUnaryServerInterceptor ...
func RecoverUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer recovery.RecoverAndLog(ctx)
		return handler(ctx, req)
	}
}
