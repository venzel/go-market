package user

type Repository interface {
	Create(user *UserEntity) (*UserEntity, error)
	List() []*UserEntity
}

type Services interface {
	Create(name string) error
	List() []*UserEntity
}
