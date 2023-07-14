package app

import (
	appTransaction "github.com/venzel/go-market/internal/market/app/transaction"
	transactionRepository "github.com/venzel/go-market/internal/market/app/transaction/infra/repository"
	appUser "github.com/venzel/go-market/internal/market/app/user"
	userRepository "github.com/venzel/go-market/internal/market/app/user/infra/repository"
	"github.com/venzel/go-market/internal/market/core/domain/transaction"
	"github.com/venzel/go-market/internal/market/core/domain/user"
	"github.com/venzel/go-market/internal/market/core/infra/database"
)

type Containers struct {
	User        user.Services
	Transaction transaction.Services
}

func (c *Containers) initUser(db database.Db) {
	userRepository := &userRepository.UserRepositoryImpl{Db: db}
	userServices := &appUser.UserServicesImpl{}
	userServices.Init(userRepository)
	c.User = userServices
}

func (c *Containers) initTransaction(db database.Db) {
	transactionRepository := &transactionRepository.TransactionRepositoryImpl{Db: db}
	transactionServices := &appTransaction.TransactionServicesImpl{}
	transactionServices.Init(transactionRepository)
	c.Transaction = transactionServices
}

func (c *Containers) Init(db database.Db) {
	c.initUser(db)
	c.initTransaction(db)
}
