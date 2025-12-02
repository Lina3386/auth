package user

import (
	"context"
)

func (s *serv) RegisterTelegramUser(ctx context.Context, telegramID int64, username string) (int64, string, error) {
	var id int64
	var token string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, token, errTx = s.userRepository.RegisterTelegramUser(ctx, telegramID, username)
		if errTx != nil {
			return errTx
		}
		return nil
	})

	if err != nil {
		return 0, "", err
	}

	return id, token, nil
}

func (s *serv) VerifyToken(ctx context.Context, token string) (int64, bool, error) {
	return s.userRepository.VerifyToken(ctx, token)
}

