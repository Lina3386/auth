package user

import (
	"context"
	"github.com/Lina3386/auth/internal/converter"
	desc "github.com/Lina3386/auth/pkg/user"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	nUser, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("%v %v %v %v %v", nUser.Id, nUser.Name, nUser.Email, nUser.Role, nUser.CreatedAt)
	return converter.ToGetResponseFromService(nUser), nil
}
