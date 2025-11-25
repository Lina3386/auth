package user

import (
	"github.com/Lina3386/auth/internal/client/db"
	"github.com/Lina3386/auth/internal/repository"
	def "github.com/Lina3386/auth/internal/service"
)

var _ def.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(userRepository repository.UserRepository, txManager db.TxManager) *serv {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
