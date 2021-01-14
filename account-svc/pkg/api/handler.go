package api

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/atyagi9006/banking-app/account-svc/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	authmgrPB "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
)

const (
	errInternal         = "Internal Error"
	errInvalidEmail     = "Invalid email"
	errInvalidPassword  = "Invalid Password"
	errEmptyFullName    = "Full Name can't be empty"
	errInvalidArgument  = "Invalid Argument"
	errInvalidRole      = "Invalid Role"
	errNoRows           = "no rows"
	errEmployeeExists   = "Employee already Exists"
	errEmployeeNotFound = "Employee not found"
)

var roles = map[string]string{
	"admin": "admin",
	"staff": "staff",
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
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new employee: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}

	if existingEmp.Email == req.Email {
		return nil, status.Error(codes.AlreadyExists, errEmployeeExists)
	}

	//register user to auth
	registerReq := authmgrPB.RegisterUserRequest{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
	regResp, err := svc.authmgrClient.RegisterUser(ctx, &registerReq)
	if err != nil {
		log.Println("Error with create new employee failed while auth reg: ", err)
		return nil, status.Error(codes.Internal, errInternal)
	}

	log.Println("user is registered successfully on auth ID:", regResp.Id)

	res, err := svc.store.CreateEmployee(ctx, svc.fromCreateEmployeeProto(req))
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}

	emp := svc.toEmployeeProto(res)
	return emp, nil
}

func (svc *AccountService) DeleteEmployee(ctx context.Context, req *pb.EmployeeIdRequest) (*pb.EmptyMessage, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}
	reqGetEmp := pb.GetEmployeeRequest{
		Id: req.Id,
	}
	getEmpResp, err := svc.GetEmployee(ctx, &reqGetEmp)
	if err != nil {
		return nil, err
	}

	id, err := strconv.ParseInt(getEmpResp.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}
	err = svc.store.DeleteEmployee(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}
	return &pb.EmptyMessage{}, nil
}

func (svc *AccountService) GetEmployee(ctx context.Context, req *pb.GetEmployeeRequest) (*pb.Employee, error) {
	var res *pb.Employee
	if req.Id == "" && req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	if !govalidator.IsEmail(req.Email) && req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidEmail)
	}

	if req.Email != "" {
		emp, err := svc.store.GetEmployeeByEmail(ctx, req.Email)
		if err != nil {

			if strings.Contains(err.Error(), errNoRows) {
				return nil, status.Error(codes.NotFound, errEmployeeNotFound)
			}
			log.Println("Error while get employee: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
		res = svc.toEmployeeProto(emp)
	}

	if req.Id != "" {
		id, err := strconv.ParseInt(req.Id, 10, 64)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
		}
		emp, err := svc.store.GetEmployee(ctx, id)
		if err != nil {
			log.Println("Error ", err)
			if strings.Contains(err.Error(), errNoRows) {
				return nil, status.Error(codes.NotFound, errEmployeeNotFound)
			}
			return nil, status.Error(codes.Internal, errInternal)
		}
		res = svc.toEmployeeProto(emp)
	}

	return res, nil
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
