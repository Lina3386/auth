package user

import (
	"context"
	desc "github.com/Lina3386/auth/pkg/user"
	"log"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	log.Printf("User deleted %v", req.GetId())
	return &desc.DeleteResponse{}, nil
}
