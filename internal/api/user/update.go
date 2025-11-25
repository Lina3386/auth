package user

import (
	"context"
	"github.com/Lina3386/auth/internal/converter"
	desc "github.com/Lina3386/auth/pkg/user"
	"log"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	err := i.userService.Update(ctx, converter.ToUserModelUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	log.Printf("User %d updated", req.GetId())
	return &desc.UpdateResponse{}, nil
}
