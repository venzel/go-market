package user

type UserEntity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserEntity(id, name string) *UserEntity {
	return &UserEntity{
		ID:   id,
		Name: name,
	}
}
