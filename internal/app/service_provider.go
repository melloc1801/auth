package app

import (
	"auth/internal/api"
	"auth/internal/closer"
	"auth/internal/config"
	"auth/internal/repository"
	"auth/internal/repository/user"
	"auth/internal/service"
	userService "auth/internal/service/user"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pg             *pgxpool.Pool
	userRepository repository.UserRepository

	userService service.UserService

	userImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GrpcConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpcConfig, err := config.NewGrpcConfig()
		if err != nil {
			log.Fatalf("failed to create grpc config")
		}
		s.grpcConfig = grpcConfig
	}

	return s.grpcConfig
}

func (s *serviceProvider) PgConfig() config.PGConfig {
	if s.pgConfig == nil {
		pgConfig, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to create pg config")
		}
		s.pgConfig = pgConfig
	}

	return s.pgConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) *pgxpool.Pool {
	if s.pg == nil {
		pool, err := pgxpool.Connect(ctx, s.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err)
		}

		closer.Add(func() error {
			pool.Close()

			return nil
		})

		s.pg = pool
	}

	return s.pg
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = user.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *api.Implementation {
	if s.userImpl == nil {
		s.userImpl = api.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
