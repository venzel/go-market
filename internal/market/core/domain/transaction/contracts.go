package transaction

type Repository interface {
	Create(user *TransactionEntity) (*TransactionEntity, error)
	List() []*TransactionEntity
}

type Services interface {
	Create(name string) error
	List() []*TransactionEntity
}
