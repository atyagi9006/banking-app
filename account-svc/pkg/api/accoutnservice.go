package api

import (
	"errors"

	"github.com/atyagi9006/banking-app/account-svc/auth"
	"github.com/atyagi9006/banking-app/account-svc/db"
	"github.com/atyagi9006/banking-app/account-svc/pkg/config"
)

type AccountService struct {
	store *db.Store

	jwtManager      *auth.JWTManager
	accessibleRoles map[string][]string
}

func NewAccountService() (*AccountService, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return nil, errors.New("config was nil")
	}
	store := db.NewStore(cfg.DBConfig)
	jwtMangager := auth.NewJWTManager(auth.SecretKey, auth.TokenDuration)
	accountService := AccountService{
		store:           store,
		jwtManager:      jwtMangager,
		accessibleRoles: accessibleRoles(),
	}
	return &accountService, nil
}
