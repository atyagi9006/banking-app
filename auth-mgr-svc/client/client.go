package client

import (
	pb "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
	grpc_mw "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// NewClient creates an instance of auth-mrg-svc client
func NewClient(addr string, interceptors ...grpc.UnaryClientInterceptor) (pb.AuthMgrServiceClient, error) {
	options := []grpc.DialOption{grpc.WithInsecure()}

	interceptor := grpc_mw.ChainUnaryClient(interceptors...)
	options = append(options, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial(addr, options...)
	if err != nil {
		return nil, err
	}

	return pb.NewAuthMgrServiceClient(conn), nil
}
