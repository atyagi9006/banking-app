package api

import (
	"errors"

	"github.com/atyagi9006/banking-app/auth-mgr-svc/db"
	"github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/auth"
	"github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/config"
)

type AuthMgrService struct {
	jwtManager auth.TokenProvider //generate token and extract token metadata
	authStore  auth.AuthStore     //store token in redis
	store      *db.Store
}

func NewAuthMgrService() (*AuthMgrService, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return nil, errors.New("config was nil")
	}

	store := db.NewStore(cfg.DBConfig)
	jwtMangager := auth.NewJWTManager(cfg.JWtConfig)
	authStore := auth.NewAuthStore(cfg.RedisConfig)
	AuthMgrService := AuthMgrService{
		store:      store,
		jwtManager: jwtMangager,
		authStore:  authStore,
	}
	return &AuthMgrService, nil
}
