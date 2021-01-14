package api

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/atyagi9006/banking-app/auth-mgr-svc/db"
	pb "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (svc *AuthMgrService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.User, error) {
	if !govalidator.IsEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, errInvalidEmail)
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidPassword)
	}

	if req.Role == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidRole)
	}

	_, err := svc.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with register new user: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}

	res, err := svc.store.CreateUser(ctx, svc.fromCreateUserProto(req))
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}

	user := svc.toUserProto(res)
	return user, nil
}
func (svc *AuthMgrService) DeleteUser(ctx context.Context, in *pb.UserIdRequest) (*pb.EmptyMessageResponse, error) {
	return nil, nil
}
func (svc *AuthMgrService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	return nil, nil
}

func (svc *AuthMgrService) fromCreateUserProto(req *pb.RegisterUserRequest) db.CreateUserParams {
	return db.CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (svc *AuthMgrService) toUserProto(user db.User) *pb.User {
	return &pb.User{
		Id:    fmt.Sprintf("%v", user.ID),
		Email: user.Email,
		Role:  user.Role,
	}
}
