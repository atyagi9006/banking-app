package api

import (
	"context"
	"log"

	pb "github.com/atyagi9006/banking-app/auth-svc/pkg/proto"
)

type AuthService struct {
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
