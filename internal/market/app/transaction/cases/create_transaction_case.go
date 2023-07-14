package cases

import "github.com/venzel/go-market/internal/market/core/domain/transaction"

type CreateTransactionCase interface {
	Execute(name string) (*transaction.TransactionEntity, error)
}

type CreateTransactionCaseImpl struct {
	TransactionRepository transaction.Repository
}

func (u *CreateTransactionCaseImpl) Execute(name string) (*transaction.TransactionEntity, error) {
	transactionEntity := transaction.NewTransactionEntity("10", name)

	transaction, err := u.TransactionRepository.Create(transactionEntity)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
