package api

import (
	"context"
	"log"

	pb "github.com/atyagi9006/banking-app/auth-svc/pkg/proto"
)

type HelloGRPCService struct {
}

//SayHello is implementing grpc-hello-world
func (svc *HelloGRPCService) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "Milgya greeting...bas rehn de....",
	}
	return &res, nil
}

//SayHellogw is implementing grpc-gw-hello-world
func (svc *HelloGRPCService) SayHellogw(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg in gw : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "Milgya greeting...bas rehn de.... gw",
	}
	return &res, nil
}
