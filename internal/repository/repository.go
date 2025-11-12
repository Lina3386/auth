package repository

import (
	"context"
	"github.com/Lina3386/auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, req *model.UserToCreate)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, req *model.UserToUpdate)
	Delete(ctx context.Context, id int64) error
}
