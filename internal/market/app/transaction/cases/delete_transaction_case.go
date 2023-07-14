package cases

import "github.com/venzel/go-market/internal/market/core/domain/transaction"

type DeleteTransactionCase interface {
	Execute() []*transaction.TransactionEntity
}

type DeleteTransactionCaseImpl struct {
	TransactionRepository transaction.Repository
}

func (u *DeleteTransactionCaseImpl) Execute() []*transaction.TransactionEntity {
	return nil
}
