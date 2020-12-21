package api

import (
	"database/sql"
	"log"

	"github.com/atyagi9006/banking-app/account-svc/db"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:P@ssw0rd@localhost:5432/simple_bank?sslmode=disable"
)

type AccountService struct {
	store *db.Store
}

func NewAccountService() *AccountService {
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(testDB)

	accountService := AccountService{
		store: store,
	}
	return &accountService
}
