package mw

import (
	"context"
	"net"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// TraceUnaryServerInterceptor ...
func TraceUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if span := opentracing.SpanFromContext(ctx); span != nil {
			traceIDToTrailer(ctx, span)
			tagRemoteAddr(ctx, span)
			tagXForwardedFor(ctx, span)
		}
		return handler(ctx, req)

	}
}

// TraceStreamServerInterceptor ...
func TraceStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if span := opentracing.SpanFromContext(ss.Context()); span != nil {
			traceIDToTrailer(ss.Context(), span)
			tagRemoteAddr(ss.Context(), span)
			tagXForwardedFor(ss.Context(), span)
		}
		return handler(srv, ss)
	}
}

func traceIDToTrailer(ctx context.Context, span opentracing.Span) {
	if sc, ok := span.Context().(jaeger.SpanContext); ok {
		trailer := metadata.Pairs(responseTraceIDHeader, sc.TraceID().String())
		_ = grpc.SetTrailer(ctx, trailer)
	}
}

func tagRemoteAddr(ctx context.Context, span opentracing.Span) {
	if p, ok := peer.FromContext(ctx); ok {
		if host, _, err := net.SplitHostPort(p.Addr.String()); err == nil {
			span.SetTag("remote.addr", host)
		}
	}
}

func tagXForwardedFor(ctx context.Context, span opentracing.Span) {
	if m, ok := metadata.FromIncomingContext(ctx); ok {
		if values := m.Get("x-forwarded-for"); len(values) > 0 {
			span.SetTag("x.forwarded.for", strings.Join(values, ","))
		}
	}
}
