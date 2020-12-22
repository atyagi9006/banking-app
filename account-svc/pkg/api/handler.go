package api

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/atyagi9006/banking-app/account-svc/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
)

const (
	errInternal        = "Internal Error"
	errInvalidEmail    = "Invalid email"
	errInvalidPassword = "Invalid Password"
	errEmptyFullName   = "Full Name can't be empty"
	errInvalidArgument = "Invalid Argument"
	errInvalidRole     = "Invalid Role"
	errNoRows          = "no rows"
	errEmployeeExists  = "Employee already Exists"
)

var roles = map[string]string{
	"admin": "admin",
	"staff": "staff",
}

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
	if !govalidator.IsEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, errInvalidEmail)
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidPassword)
	}

	if req.FullName == "" {
		return nil, status.Error(codes.InvalidArgument, errEmptyFullName)
	}

	if _, ok := roles[req.Role]; !ok {
		return nil, status.Error(codes.InvalidArgument, errInvalidRole)
	}

	existingEmp, err := svc.store.GetEmployeeByEmail(ctx, req.Email)
	if err != nil {
		log.Println("Error ", err)
		if !strings.Contains(err.Error(), errNoRows) {
			return nil, status.Error(codes.Internal, errInternal)
		}
	}

	if existingEmp.Email == req.Email {
		return nil, status.Error(codes.AlreadyExists, errEmployeeExists)
	}

	res, err := svc.store.CreateEmployee(ctx, svc.fromCreateEmployeeProto(req))
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
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
		FullName: emp.FullName,
		Role:     emp.Role,
	}
}
