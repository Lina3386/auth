package service

import (
	"context"
	"github.com/Lina3386/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, req *model.UserToCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, req *model.UserToUpdate) error
	RegisterTelegramUser(ctx context.Context, telegramID int64, username string) (int64, string, error)
	VerifyToken(ctx context.Context, token string) (int64, bool, error)
}
