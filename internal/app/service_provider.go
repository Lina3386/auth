package app

import (
	"context"
	"github.com/Lina3386/auth/internal/api/user"
	"github.com/Lina3386/auth/internal/client/db"
	"github.com/Lina3386/auth/internal/client/db/pg"
	"github.com/Lina3386/auth/internal/client/transaction"
	"github.com/Lina3386/auth/internal/closer"
	"github.com/Lina3386/auth/internal/config"
	"github.com/Lina3386/auth/internal/config/env"
	"github.com/Lina3386/auth/internal/repository"
	userRepo "github.com/Lina3386/auth/internal/repository/user"
	"github.com/Lina3386/auth/internal/service"
	userService "github.com/Lina3386/auth/internal/service/user"
	"log"
)

type ServiceProvider struct {
	pgCongfig  config.PGConfig
	grpcConfig config.GRPCConfig

	dbclient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository
	userService    service.UserService

	userImpl *user.Implementation
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) PGConfig() config.PGConfig {
	if s.pgCongfig == nil {
		pgConfig, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgCongfig = pgConfig
	}

	return s.pgCongfig
}

func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpcConfig, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}
		s.grpcConfig = grpcConfig
	}
	return s.grpcConfig
}

func (s *ServiceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbclient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to get db client: %v", err)
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("oing error: %s", err.Error())
		}

		closer.Add(cl.Close)
		s.dbclient = cl
	}
	return s.dbclient
}

func (s *ServiceProvider) PgPool(ctx context.Context) db.Client {
	if s.dbclient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		closer.Add(cl.Close)
		s.dbclient = cl
	}
	return s.dbclient
}

func (s *ServiceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *ServiceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewRepository(s.PgPool(ctx))
	}
	return s.userRepository
}

func (s *ServiceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx), s.TxManager(ctx))
	}
	return s.userService
}

func (s *ServiceProvider) UserImplementation(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}
	return s.userImpl
}
