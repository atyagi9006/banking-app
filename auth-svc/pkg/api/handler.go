package api

import (
	"context"
	"log"

	authmgrClient "github.com/atyagi9006/banking-app/auth-mgr-svc/client"
	authmgrPB "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
	pb "github.com/atyagi9006/banking-app/auth-svc/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	authmgrClient authmgrPB.AuthMgrServiceClient
}

func NewAuthService() *AuthService {

	authmgrClient, err := authmgrClient.NewClient(":7781")
	if err != nil {
		log.Fatal("cannot connect to account-svc", err)
	}
	authService := AuthService{
		authmgrClient: authmgrClient,
	}
	return &authService
}

func (svc *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	md := metadata.Pairs("mode", "internal")
	newCtx := metadata.NewOutgoingContext(ctx, md)
	tokenReq := authmgrPB.GenerateTokenRequest{
		Email:    req.Username,
		Password: req.Password,
	}
	tokenres, err := svc.authmgrClient.GenerateToken(newCtx, &tokenReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{
		AccessToken:  tokenres.AccessToken,
		RefreshToken: tokenres.RefreshToken,
	}
	return res, nil
}
