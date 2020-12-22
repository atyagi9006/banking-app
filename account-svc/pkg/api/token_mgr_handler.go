package api

import (
	"context"

	"github.com/asaskevich/govalidator"
	"github.com/atyagi9006/banking-app/account-svc/db"
	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (svc *AccountService) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	if !govalidator.IsEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, errInvalidEmail)
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidPassword)
	}

	employee := db.BankEmployee{
		Email:    req.Email,
		Password: req.Password,
	}
	token, err := svc.jwtManager.Generate(&employee)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	res := pb.GenerateTokenResponse{Token: token}
	return &res, nil
}
