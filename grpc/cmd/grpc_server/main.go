package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Lina3386/auth/grpc/proto"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserServiceServer
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("======= Create User =======")
	log.Printf("Name: %s", req.GetName())
	log.Printf("Email: %s", req.GetEmail())
	log.Printf("Password: %s", req.GetPassword())
	log.Printf("Password_confirm: %s", req.GetPasswordConfirm())
	log.Printf("Role: %s", req.GetRole().String())

	return *desc.CreateResponse{
		Id: req.GetId(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("======= Get User =======")
	log.Printf("Id: %d", req.GetId())

	now := timestamppb.Now()

	return *desc.GetResponse{
		Id:        req.GetId(),
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Role:      pb.ROLE_USER,
		CreatedAt: now,
		UpdateAt:  now,
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	log.Printf("======= Update User =======")
	log.Printf("Id: %d", req.GetId())
	log.Printf("New name: %s", req.GetName().GetValue())
	log.Printf("New email: %s", req.GetEmail().GetValue())
	log.Printf("Role: %s", req.GetRole().String())

	return *desc.UpdateResponse{
		Empty: emptypb.Empty{},
	}, nil
}
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	log.Printf("======= Delete User =======")
	log.Printf("Id: %d", req.GetId())

	return *desc.DeleteResponse{
		Empty: emptypb.Empty{},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nill {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
	}

}
