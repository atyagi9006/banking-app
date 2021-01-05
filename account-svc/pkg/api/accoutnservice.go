package api

import (
	"context"
	"errors"
	"log"

	"github.com/atyagi9006/banking-app/account-svc/auth"
	"github.com/atyagi9006/banking-app/account-svc/db"
	"github.com/atyagi9006/banking-app/account-svc/pkg/config"
	"github.com/atyagi9006/opa-authz/opa"
)

const (
	policyName = "rego/authz.rego"
)

type AccountService struct {
	store           *db.Store
	config          *config.SVCConfig
	opaClient       *opa.Client
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
	opaClient, err := opa.NewClient(cfg.OPAConfig.Endpoint)
	if err != nil {
		log.Fatal("opa is not started")
	}
	//create policy in opa for authorization
	opaClient.CreatePolicyFromFile(context.Background(), policyName)

	accountService := AccountService{
		store:           store,
		config:          cfg,
		jwtManager:      jwtMangager,
		opaClient:       opaClient,
		accessibleRoles: accessibleRoles(),
	}
	return &accountService, nil
}
