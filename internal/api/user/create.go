package user

import (
	"context"
	"github.com/Lina3386/auth/internal/converter"
	desc "github.com/Lina3386/auth/pkg/user"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserModelCreateFromDesc(req))
	if err != nil {
		return nil, err
	}
	log.Printf("insered user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
