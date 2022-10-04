package main

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"lib/core/v1"
	"lib/core/v1/logger"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := core.LoadConfig("./.cfg/k8s")
	if err != nil {
		logger.Fatal(ctx, err)
	}

	logger.SetLevel(logger.FromConfig(viper.GetString("env.log_level")))

	a := core.New(ctx)
	a.WithUnaryMW(exampleOfMW)

	if err = a.Run(exampleOfSvc()); err != nil {
		logger.Fatalf(ctx, "can't run app: %v", err)
	}
}

func exampleOfMW(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(ctx, req)
}

func exampleOfSvc() core.IService {
	return core.RegisterFunc(func(_ grpc.ServiceRegistrar) {
		// do something
	})
}
