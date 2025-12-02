package repository

import (
	"context"
	"github.com/Lina3386/auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, req *model.UserToCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, req *model.UserToUpdate) error
	Delete(ctx context.Context, id int64) error
	RegisterTelegramUser(ctx context.Context, telegramID int64, username string) (int64, string, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*model.User, error)
	VerifyToken(ctx context.Context, token string) (int64, bool, error)
}
