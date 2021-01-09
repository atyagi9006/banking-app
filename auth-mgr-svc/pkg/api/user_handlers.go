package api

import (
	"context"

	pb "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
)

func (svc *AuthMgrService) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.User, error) {
	return nil, nil
}
func (svc *AuthMgrService) DeleteUser(ctx context.Context, in *pb.UserIdRequest) (*pb.EmptyMessage, error) {
	return nil, nil
}
func (svc *AuthMgrService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	return nil, nil
}
