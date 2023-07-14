package cases

import "github.com/venzel/go-market/internal/market/core/domain/user"

type CreateUserCase interface {
	Execute(name string) (*user.UserEntity, error)
}

type CreateUserCaseImpl struct {
	UserRepository user.Repository
}

func (u *CreateUserCaseImpl) Execute(name string) (*user.UserEntity, error) {
	userEntity := user.NewUserEntity("10", name)

	user, err := u.UserRepository.Create(userEntity)

	if err != nil {
		return nil, err
	}

	return user, nil
}
