package repository

import (
	"github.com/venzel/go-market/internal/market/core/domain/transaction"
	"github.com/venzel/go-market/internal/market/core/infra/database"
)

type TransactionRepositoryImpl struct {
	Db database.Db
}

func (r *TransactionRepositoryImpl) Create(t *transaction.TransactionEntity) (*transaction.TransactionEntity, error) {
	transactionGorm := database.NewTransactionGorm(t.ID, t.Name)

	err := r.Db.SaveTransaction(transactionGorm)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *TransactionRepositoryImpl) List() []*transaction.TransactionEntity {
	transactionsGorm := r.Db.FindTransactions()

	transactions := []*transaction.TransactionEntity{}

	for _, transactionGorm := range transactionsGorm {
		transaction := transaction.NewTransactionEntity(transactionGorm.ID, transactionGorm.Name)
		transactions = append(transactions, transaction)
	}

	return transactions
}
