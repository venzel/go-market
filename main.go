package main

import (
	"fmt"

	"github.com/venzel/go-market/internal/market/app"
	"github.com/venzel/go-market/internal/market/core/infra/database"
)

func main() {
	db := &database.Gorm{}
	db.InitDb()

	containers := &app.Containers{}
	containers.Init(db)

	containers.User.Create("Venzel")
	containers.Transaction.Create("Bitcoin")

	for _, user := range containers.User.List() {
		fmt.Println(user.Name)
	}

	for _, transaction := range containers.Transaction.List() {
		fmt.Println(transaction.Name)
	}
}
