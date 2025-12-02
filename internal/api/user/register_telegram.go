package user

import (
	"context"
	"log"

	desc "github.com/Lina3386/auth/pkg/user"
)

func (i *Implementation) RegisterTelegramUser(ctx context.Context, req *desc.RegisterTelegramUserRequest) (*desc.RegisterTelegramUserResponse, error) {
	userID, token, err := i.userService.RegisterTelegramUser(ctx, req.TelegramId, req.Username)
	if err != nil {
		log.Printf("Failed to register telegram user: %v", err)
		return nil, err
	}

	log.Printf("Registered telegram user: ID=%d, TelegramID=%d, Username=%s", userID, req.TelegramId, req.Username)

	return &desc.RegisterTelegramUserResponse{
		UserId: userID,
		Token:  token,
	}, nil
}

func (i *Implementation) VerifyToken(ctx context.Context, req *desc.VerifyTokenRequest) (*desc.VerifyTokenResponse, error) {
	userID, valid, err := i.userService.VerifyToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	return &desc.VerifyTokenResponse{
		UserId: userID,
		Valid:  valid,
	}, nil
}

