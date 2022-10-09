package main

import (
	"context"
	"fmt"

	core "github.com/gladilindv/appcore"
	"github.com/gladilindv/appcore/logger"
	"google.golang.org/grpc"
)

const cfgSvcName = "env.svc.name"

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := core.LoadConfig()
	if err != nil {
		logger.Fatal(ctx, err)
	}

	logger.InitFromConfig(core.LogLevel())

	a := core.New(ctx)
	a.WithUnaryMW(exampleOfMW)

	svcName := core.ConfigValue(cfgSvcName).String()
	if err = a.Run(exampleOfSvc(svcName)); err != nil {
		logger.Fatalf(ctx, "can't run app: %v", err)
	}
}

func exampleOfMW(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(ctx, req)
}

func exampleOfSvc(svcName string) core.IService {
	return core.RegisterFunc(func(_ grpc.ServiceRegistrar) {
		// do something
		fmt.Println("service", svcName, "registered")
	})
}
