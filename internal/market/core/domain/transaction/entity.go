package transaction

type TransactionEntity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewTransactionEntity(id, name string) *TransactionEntity {
	return &TransactionEntity{
		ID:   id,
		Name: name,
	}
}
