package main

import (
	"context"
	"fmt"
	"github.com/Lina3386/auth/internal/config"
	"github.com/Lina3386/auth/internal/store"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	desc "github.com/Lina3386/auth/pkg/user"
)

type server struct {
	desc.UnimplementedUserAPIServer
	store *store.Store
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create user: name=%s, email=%s, role=%s", req.GetName(), req.GetEmail(), req.GetRole())

	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, fmt.Errorf("password does not match")
	}

	id, err := s.store.CreateUser(ctx, req.GetName(), req.GetEmail(), req.GetPassword(), int32(req.GetRole()))
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, err
	}

	log.Printf("Create user: id=%d", id)
	return &desc.CreateResponse{Id: id}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get User: id=%d", req.GetId())

	user, err := s.store.GetUser(ctx, req.GetId())
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return nil, err
	}

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdateAt:  timestamppb.New(user.UpdatedAt),
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	log.Printf("Update user: id=%d", req.GetId())

	var name, email *string
	if req.GetName() != nil {
		v := req.GetName().GetValue()
		name = &v
	}

	if req.GetEmail() != nil {
		v := req.GetEmail().GetValue()
		email = &v
	}

	if err := s.store.UpdateUser(ctx, req.GetId(), name, email, int32(req.GetRole())); err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, err
	}

	log.Printf("User %d updated", req.GetId())

	return &desc.UpdateResponse{
		Empty: &emptypb.Empty{},
	}, nil
}
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	log.Printf("Delete user: id=%d", req.GetId())

	if err := s.store.DeleteUser(ctx, req.GetId()); err != nil {
		log.Printf("Failed to delete user: %v", err)
		return nil, err
	}

	log.Printf("User %d deleted", req.GetId())
	return &desc.DeleteResponse{
		Empty: &emptypb.Empty{},
	}, nil
}

func main() {
	grpcCfg, pgCfg, err := config.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	pool, err := pgxpool.Connect(context.Background(), pgCfg.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	st := store.New(pool)

	lis, err := net.Listen("tcp", grpcCfg.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIServer(s, &server{store: st})

	log.Printf("Server listening at %v", grpcCfg.Address())
	log.Fatal(s.Serve(lis))
}
