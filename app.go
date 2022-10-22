package core

import (
	"context"
	"fmt"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

// IService is the interface for all services
type IService interface {
	Register(srv grpc.ServiceRegistrar)
}

// RegisterFunc ...
type RegisterFunc func(srv grpc.ServiceRegistrar)

// Register default implementation
func (f RegisterFunc) Register(srv grpc.ServiceRegistrar) {
	f(srv)
}

type application struct {
	opts []grpc.ServerOption

	grpcPort uint32
	grpc     net.Listener
}

// New creates core application
func New(ctx context.Context, grpcPort uint32) *application {
	a := &application{
		grpcPort: grpcPort,
	}
	a.opts = make([]grpc.ServerOption, 0)

	return a
}

// Run starts core application
func (a *application) Run(services ...IService) error {

	a.listenGRPC(a.grpcPort)

	srv := a.initGRPC()
	for _, service := range services {
		service.Register(srv)
	}

	err := srv.Serve(a.grpc)
	if err != nil {
		panic(err)
	}

	return nil
}

func (a *application) listenGRPC(port uint32) {
	grpc, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	a.grpc = grpc
}

func (a *application) initGRPC() *grpc.Server {

	//opts := []grpc.ServerOption{
	//	grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	//	grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	//
	//	grpc.ChainUnaryInterceptor(mw.TraceUnaryServerInterceptor()),
	//	grpc.ChainStreamInterceptor(mw.TraceStreamServerInterceptor()),
	//}
	opts := a.initMW()
	opts = append(opts, a.opts...)

	grpcServer := grpc.NewServer(
		opts...,
	)
	grpc_prometheus.Register(grpcServer)

	return grpcServer
}

// nolint:unused
func (a *application) initHTTP() error {
	//mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}))
	//opts := []grpc.DialOption{
	//	grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50000000)),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//}
	//err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	//if err != nil {
	//	panic(err)
	//}
	return nil
}

// WithUnaryMW adds unary interceptor
func (a *application) WithUnaryMW(unary grpc.UnaryServerInterceptor) {
	a.opts = append(a.opts, grpc.ChainUnaryInterceptor(unary))
}
