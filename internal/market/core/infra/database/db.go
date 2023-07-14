package database

type Db interface {
	InitDb()

	SaveUser(user *UserGorm) error
	FindUsers() []*UserGorm

	SaveTransaction(user *TransactionGorm) error
	FindTransactions() []*TransactionGorm
}

type UserGorm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TransactionGorm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Gorm struct {
	UsersGorm        []*UserGorm
	TransactionsGorm []*TransactionGorm
}

func NewUserGorm(id, name string) *UserGorm {
	return &UserGorm{id, name}
}

func NewTransactionGorm(id, name string) *TransactionGorm {
	return &TransactionGorm{id, name}
}

func (g *Gorm) InitDb() {
	users := []*UserGorm{}

	tiago := &UserGorm{"1", "Tiago"}
	marcos := &UserGorm{"2", "Marcos"}

	g.UsersGorm = append(users, tiago, marcos)
}

func (g *Gorm) SaveUser(user *UserGorm) error {
	userGorm := &UserGorm{user.ID, user.Name}

	g.UsersGorm = append(g.UsersGorm, userGorm)

	return nil
}

func (g *Gorm) FindUsers() []*UserGorm {
	return g.UsersGorm
}

func (g *Gorm) SaveTransaction(t *TransactionGorm) error {
	transactionGorm := &TransactionGorm{t.ID, t.Name}

	g.TransactionsGorm = append(g.TransactionsGorm, transactionGorm)

	return nil
}

func (g *Gorm) FindTransactions() []*TransactionGorm {
	return g.TransactionsGorm
}
