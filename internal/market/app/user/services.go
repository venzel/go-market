package user

import (
	"github.com/venzel/go-market/internal/market/app/user/cases"
	"github.com/venzel/go-market/internal/market/core/domain/user"
)

type UserServicesImpl struct {
	createUseCase cases.CreateUserCase
	listUseCase   cases.ListUsersCase
	deleteUseCase cases.DeleteUserCase
}

func (u *UserServicesImpl) Init(userRepository user.Repository) {
	u.createUseCase = &cases.CreateUserCaseImpl{UserRepository: userRepository}
	u.listUseCase = &cases.ListUsersCaseImpl{UserRepository: userRepository}
	u.deleteUseCase = &cases.DeleteUserCaseImpl{UserRepository: userRepository}
}

func (u *UserServicesImpl) Create(name string) error {
	_, err := u.createUseCase.Execute(name)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserServicesImpl) List() []*user.UserEntity {
	return u.listUseCase.Execute()
}
