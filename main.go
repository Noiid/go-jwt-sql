package main

import (
	"CaseMajoo/api"
	"CaseMajoo/repository"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/case_majoo")
	if err != nil {
		return nil, err
	}

	return db, nil
}
func main() {
	db, err := connect()
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	usersRepo := repository.NewUserRepository(db)
	merchantRepo := repository.NewMerchantRepository(db)
	transactionRepo := repository.NewTransactionsRepository(db)
	mainAPI := api.NewApi(usersRepo, merchantRepo, transactionRepo)
	mainAPI.Start()

}
