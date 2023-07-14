package transaction

import (
	"github.com/venzel/go-market/internal/market/app/transaction/cases"
	"github.com/venzel/go-market/internal/market/core/domain/transaction"
)

type TransactionServicesImpl struct {
	createUseCase cases.CreateTransactionCase
	listUseCase   cases.ListTransactionsCase
	deleteUseCase cases.DeleteTransactionCase
}

func (u *TransactionServicesImpl) Init(transactionRepository transaction.Repository) {
	u.createUseCase = &cases.CreateTransactionCaseImpl{TransactionRepository: transactionRepository}
	u.listUseCase = &cases.ListTransactionsCaseImpl{TransactionRepository: transactionRepository}
	u.deleteUseCase = &cases.DeleteTransactionCaseImpl{TransactionRepository: transactionRepository}
}

func (u *TransactionServicesImpl) Create(name string) error {
	_, err := u.createUseCase.Execute(name)

	if err != nil {
		return err
	}

	return nil
}

func (u *TransactionServicesImpl) List() []*transaction.TransactionEntity {
	return u.listUseCase.Execute()
}
