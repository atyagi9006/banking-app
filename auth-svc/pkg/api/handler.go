package api

import (
	"context"
	"log"

	accountClient "github.com/atyagi9006/banking-app/account-svc/client"
	accountpb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	pb "github.com/atyagi9006/banking-app/auth-svc/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	accountClient accountpb.AccountServiceClient
}

func NewAuthService() *AuthService {

	accountServiceClient, err := accountClient.NewClient(":7777")
	if err != nil {
		log.Fatal("cannot connect to account-svc", err)
	}
	authService := AuthService{
		accountClient: accountServiceClient,
	}
	return &authService
}

//SayHello is implementing grpc-hello-world
func (svc *AuthService) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "Hello auth grpc....",
	}
	return &res, nil
}

//SayHellogw is implementing grpc-gw-hello-world
func (svc *AuthService) SayHellogw(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg in gw : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "hello auth GRPC.... gw",
	}
	return &res, nil
}

func (svc *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	md := metadata.Pairs("mode", "internal")
	newCtx := metadata.NewOutgoingContext(ctx, md)
	tokenReq := accountpb.GenerateTokenRequest{
		Email:    req.Username,
		Password: req.Password,
	}
	token, err := svc.accountClient.GenerateToken(newC`tx, &tokenReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token.Token}
	return res, nil
}
