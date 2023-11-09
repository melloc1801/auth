package app

import (
	"auth/internal/closer"
	"auth/internal/config"
	desc "auth/pkg/user_v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	envPathName = "local.env"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGrpcServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(envPathName)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer()

	reflection.Register(a.grpcServer)
	desc.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *App) runGrpcServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GrpcConfig().Address())

	port := a.serviceProvider.GrpcConfig().Address()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil
	}

	fmt.Println("Server has been started")
	err = a.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
