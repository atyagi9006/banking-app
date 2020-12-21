package api

import (
	"context"
	"log"

	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
)

type AccountService struct {
}

//SayHello is implementing grpc-hello-world
func (svc *AccountService) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "AccountService hello.",
	}
	return &res, nil
}

//SayHellogw is implementing grpc-gw-hello-world
func (svc *AccountService) SayHellogw(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg in gw : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "Hello GRPC.... gw",
	}
	return &res, nil
}
