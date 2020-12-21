package api

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/atyagi9006/banking-app/account-svc/db"

	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
)

const (
	internalError = "Internal Error"
)

//SayHello is implementing grpc hello
func (svc *AccountService) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "AccountService hello.",
	}
	return &res, nil
}

//SayHellogw is implementing grpc -gw- hello
func (svc *AccountService) SayHellogw(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Received msg in gw : %s in request \n", in.Greeting)
	res := pb.PingMessage{
		Greeting: "Hello GRPC.... gw",
	}
	return &res, nil
}

func (svc *AccountService) CreateBankEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.Employee, error) {
	res, err := svc.store.CreateEmployee(ctx, svc.fromCreateEmployeeProto(req))
	if err != nil {
		return nil, status.Error(codes.Internal, internalError)
	}

	emp := svc.toEmployeeProto(res)
	return emp, nil

}

func (svc *AccountService) fromCreateEmployeeProto(req *pb.CreateEmployeeRequest) db.CreateEmployeeParams {
	return db.CreateEmployeeParams{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
		Role:     req.Role,
	}
}

func (svc *AccountService) toEmployeeProto(emp db.BankEmployee) *pb.Employee {
	return &pb.Employee{
		Id:       fmt.Sprintf("%v", emp.ID),
		Email:    emp.Email,
		Password: emp.Password,
		FullName: emp.FullName,
		Role:     emp.Role,
	}
}
