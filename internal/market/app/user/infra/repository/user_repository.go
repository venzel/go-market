package repository

import (
	"github.com/venzel/go-market/internal/market/core/domain/user"
	"github.com/venzel/go-market/internal/market/core/infra/database"
)

type UserRepositoryImpl struct {
	Db database.Db
}

func (r *UserRepositoryImpl) Create(user *user.UserEntity) (*user.UserEntity, error) {
	userGorm := database.NewUserGorm(user.ID, user.Name)

	err := r.Db.SaveUser(userGorm)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) List() []*user.UserEntity {
	usersGorm := r.Db.FindUsers()

	users := []*user.UserEntity{}

	for _, userGorm := range usersGorm {
		user := user.NewUserEntity(userGorm.ID, userGorm.Name)
		users = append(users, user)
	}

	return users
}
