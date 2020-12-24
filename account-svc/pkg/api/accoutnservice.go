package api

import (
	"database/sql"
	"log"

	"github.com/atyagi9006/banking-app/account-svc/auth"
	"github.com/atyagi9006/banking-app/account-svc/db"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:P@ssw0rd@localhost:5432/my_bank?sslmode=disable"
)

type AccountService struct {
	store *db.Store

	jwtManager      *auth.JWTManager
	accessibleRoles map[string][]string
}

func NewAccountService() *AccountService {
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(testDB)
	jwtMangager := auth.NewJWTManager(auth.SecretKey, auth.TokenDuration)
	accountService := AccountService{
		store:           store,
		jwtManager:      jwtMangager,
		accessibleRoles: accessibleRoles(),
	}
	return &accountService
}
