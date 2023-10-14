package main

import (
	"fmt"

	"github.com/agusheryanto182/go-wangkas/config"
	"github.com/agusheryanto182/go-wangkas/handler"
	"github.com/agusheryanto182/go-wangkas/router"
	"github.com/agusheryanto182/go-wangkas/transaction"
	"github.com/agusheryanto182/go-wangkas/user"
	handler2 "github.com/agusheryanto182/go-wangkas/web/handler"
)

func main() {
	db := config.NewDB()

	sessionRepository := user.NewRepository(db)
	sessionService := user.NewService(sessionRepository)
	sessionWebHandler := handler2.NewSessionHandler(sessionService)

	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	transactionWebHandler := handler2.NewTransactionsHandler2(transactionService)

	r := router.NewRouter(transactionHandler, transactionWebHandler, sessionWebHandler)

	err := r.Run()
	if err != nil {
		fmt.Println("can not run")
	}
}
