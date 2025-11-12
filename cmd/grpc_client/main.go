package main

import (
	"context"
	desc "github.com/Lina3386/auth/pkg/user"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := desc.NewUserAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	CrResp, err := client.Create(ctx, &desc.CreateRequest{
		Name:            "Polina",
		Email:           "polinazxz123@gmail.com",
		Password:        "11111111",
		PasswordConfirm: "11111111",
		Role:            desc.Role_USER,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CreateResponse: new user Id = %d", CrResp.GetId())

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getResp, err := client.Get(ctx, &desc.GetRequest{Id: CrResp.GetId()})
	if err != nil {
		log.Fatal("Get: %v", err)
	}
	log.Printf("GetResponse: id=%d, name=%s, email=%s, role=%s", getResp.GetId(), getResp.GetName(), getResp.GetEmail(), getResp.GetRole())

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	UpResp, err := client.Update(ctx, &desc.UpdateRequest{
		Id:   CrResp.GetId(),
		Role: desc.Role_ADMIN,
	})
	if err != nil {
		log.Fatal("Update: %v", err)
	}

	log.Printf("UpdateResponse: succes=%v", UpResp.GetEmpty() != nil)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	DelResp, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: CrResp.GetId(),
	})
	if err != nil {
		log.Fatal("Delete: %v", err)
	}

	log.Printf("DeleteResponse: %d\n", DelResp.GetEmpty() != nil)
}
