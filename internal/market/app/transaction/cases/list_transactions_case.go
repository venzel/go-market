package cases

import "github.com/venzel/go-market/internal/market/core/domain/transaction"

type ListTransactionsCase interface {
	Execute() []*transaction.TransactionEntity
}

type ListTransactionsCaseImpl struct {
	TransactionRepository transaction.Repository
}

func (u *ListTransactionsCaseImpl) Execute() []*transaction.TransactionEntity {
	return u.TransactionRepository.List()
}
