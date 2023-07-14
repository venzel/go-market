package cases

import "github.com/venzel/go-market/internal/market/core/domain/user"

type ListUsersCase interface {
	Execute() []*user.UserEntity
}

type ListUsersCaseImpl struct {
	UserRepository user.Repository
}

func (u *ListUsersCaseImpl) Execute() []*user.UserEntity {
	return u.UserRepository.List()
}
