package user

import (
	"github.com/Lina3386/auth/internal/service"
	desc "github.com/Lina3386/auth/pkg/user"
)

type Implementation struct {
	desc.UnimplementedUserAPIServer
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{userService: userService}
}
