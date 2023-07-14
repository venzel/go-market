package cases

import "github.com/venzel/go-market/internal/market/core/domain/user"

type DeleteUserCase interface {
	Execute() []*user.UserEntity
}

type DeleteUserCaseImpl struct {
	UserRepository user.Repository
}

func (u *DeleteUserCaseImpl) Execute() []*user.UserEntity {
	return nil
}
