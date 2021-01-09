package api

import (
	"context"
	"log"
	"strings"

	"github.com/asaskevich/govalidator"
	pb "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errInvaildUserNamePassword = "Incorrect username or password"
	errInternal                = "Internal Error"
	errInvalidEmail            = "Invalid email"
	errInvalidPassword         = "Invalid Password"
	errInvalidArgument         = "Invalid Argument"
	errInvalidRole             = "Invalid Role"
	errNoRows                  = "no rows"
	errUserExists              = "User already Exists"
	errUserNotFound            = "User not found"
)

func (svc *AuthMgrService) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	if !govalidator.IsEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, errInvalidEmail)
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidPassword)
	}

	emp, err := svc.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Println("Error ", err)
		if strings.Contains(err.Error(), errNoRows) {
			return nil, status.Error(codes.NotFound, errUserNotFound)
		}
		return nil, status.Error(codes.Internal, errInternal)
	}

	if req.Password != emp.Password {
		return nil, status.Error(codes.InvalidArgument, errInvaildUserNamePassword)
	}

	token, err := svc.jwtManager.Generate(&emp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	res := pb.GenerateTokenResponse{Token: token}
	return &res, nil
}
