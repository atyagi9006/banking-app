package api

import (
	"context"
	"log"

	accountClient "github.com/atyagi9006/banking-app/account-svc/client"
	accountpb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	pb "github.com/atyagi9006/banking-app/auth-mrg-svc/pkg/proto"
)

type AuthMgrService struct {
	accountClient accountpb.AccountServiceClient
}

func NewAuthMgrService() *AuthMgrService {

	accountServiceClient, err := accountClient.NewClient(":7777")
	if err != nil {
		log.Fatal("cannot connect to account-svc", err)
	}
	AuthMgrService := AuthMgrService{
		accountClient: accountServiceClient,
	}
	return &AuthMgrService
}

func (svc *AuthMgrService) GenerateToken(ctx context.Context, in *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	return nil, nil
}
