package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcioecom/wallet-core/internal/database"
	"github.com/marcioecom/wallet-core/internal/event"
	"github.com/marcioecom/wallet-core/internal/usecase/createaccount"
	"github.com/marcioecom/wallet-core/internal/usecase/createclient"
	"github.com/marcioecom/wallet-core/internal/usecase/createtransaction"
	"github.com/marcioecom/wallet-core/internal/web"
	"github.com/marcioecom/wallet-core/pkg/events"
	"github.com/marcioecom/wallet-core/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.New(ctx, db)
	uow.Register("account", func(tx *sql.Tx) any {
		return database.NewAccountDB(db)
	})
	uow.Register("transaction", func(tx *sql.Tx) any {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	clientHandler := web.NewClientHandler(*createClientUseCase)
	accountHandler := web.NewAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewTransactionHandler(*createTransactionUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /clients", clientHandler.CreateClient)
	mux.HandleFunc("POST /accounts", accountHandler.CreateAccount)
	mux.HandleFunc("POST /transactions", transactionHandler.CreateTransaction)

	log.Println("Listening on PORT", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Failed to listen and serve", err)
	}
}
