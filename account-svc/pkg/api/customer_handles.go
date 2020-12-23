package api

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/atyagi9006/banking-app/account-svc/db"
	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errEmptyAddress     = "Addess can't be empty"
	errInvalidKycType   = "Invalid KycType"
	errEmptyKycID       = "KYC ID Name can't be empty"
	errCustomerExists   = "Customer already Exists"
	errCustomerNotFound = "Customer not found"
	errKycAlreadyExists = "Customer Kyc already Exists"
)

var kycType = map[string]string{
	"Pan Card":    "Pan Card",
	"Aadhar Card": "Aadhar Card",
	"Voter Card":  "Voter Card",
	"Password":    "Password",
}

func (svc *AccountService) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.Customer, error) {
	if !govalidator.IsEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, errInvalidEmail)
	}

	if req.FullName == "" {
		return nil, status.Error(codes.InvalidArgument, errEmptyFullName)
	}

	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, errEmptyAddress)
	}

	if _, ok := kycType[req.KycType]; !ok {
		return nil, status.Error(codes.InvalidArgument, errInvalidKycType)
	}

	if req.KycId == "" {
		return nil, status.Error(codes.InvalidArgument, errEmptyKycID)
	}

	existingCus, err := svc.store.GetCustomerByEmail(ctx, req.Email)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new Customer: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}

	if existingCus.FullName == req.FullName {
		return nil, status.Error(codes.AlreadyExists, errCustomerExists)
	}

	res, err := svc.store.CreateCustomer(ctx, svc.fromCreateCustomerProto(req))
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}

	cus := svc.toCustomerProto(res)
	return cus, nil
}

func (svc *AccountService) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.Customer, error) {

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	if _, ok := kycType[req.KycType]; !ok {
		return nil, status.Error(codes.InvalidArgument, errInvalidKycType)
	}
	if req.KycId == "" {
		return nil, status.Error(codes.InvalidArgument, errEmptyKycID)
	}

	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	existingCus, err := svc.store.GetCustomer(ctx, id)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new Customer: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}

	if existingCus.KycID == req.KycId {
		return nil, status.Error(codes.AlreadyExists, errKycAlreadyExists)
	}
	args := db.UpdateCustomerKYCParams{
		ID:      id,
		KycType: req.KycType,
		KycID:   req.KycId,
	}
	res, err := svc.store.UpdateCustomerKYC(ctx, args)
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}
	cus := svc.toCustomerProto(res)
	return cus, nil
}

func (svc *AccountService) DeleteCustomer(ctx context.Context, req *pb.CustomerIdRequest) (*pb.EmptyMessage, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}
	reqGetCus := pb.GetCustomerRequest{
		Id: req.Id,
	}
	getCusResp, err := svc.GetCustomer(ctx, &reqGetCus)
	if err != nil {
		return nil, err
	}

	id, err := strconv.ParseInt(getCusResp.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}
	err = svc.store.DeleteCustomer(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}
	return &pb.EmptyMessage{}, nil
}

func (svc *AccountService) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.Customer, error) {
	var res *pb.Customer
	if req.Id == "" && req.Email == "" && req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	if !govalidator.IsEmail(req.Email) && req.Id == "" && req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidEmail)
	}

	if req.Email != "" {
		cus, err := svc.store.GetCustomerByEmail(ctx, req.Email)
		if err != nil {

			if strings.Contains(err.Error(), errNoRows) {
				return nil, status.Error(codes.NotFound, errCustomerNotFound)
			}
			log.Println("Error while get Customer: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
		res = svc.toCustomerProto(cus)
	}

	if req.Id != "" {
		id, err := strconv.ParseInt(req.Id, 10, 64)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
		}
		cus, err := svc.store.GetCustomer(ctx, id)
		if err != nil {
			log.Println("Error ", err)
			if strings.Contains(err.Error(), errNoRows) {
				return nil, status.Error(codes.NotFound, errCustomerNotFound)
			}
			return nil, status.Error(codes.Internal, errInternal)
		}
		res = svc.toCustomerProto(cus)
	}

	if req.Name != "" {
		cus, err := svc.store.GetCustomerByFullName(ctx, req.Name)
		if err != nil {

			if strings.Contains(err.Error(), errNoRows) {
				return nil, status.Error(codes.NotFound, errCustomerNotFound)
			}
			log.Println("Error while get Customer: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
		res = svc.toCustomerProto(cus)
	}

	return res, nil
}

func (svc *AccountService) fromCreateCustomerProto(req *pb.CreateCustomerRequest) db.CreateCustomerParams {
	return db.CreateCustomerParams{
		Email:    req.Email,
		FullName: req.FullName,
		Address:  req.Address,
		KycType:  req.KycType,
		KycID:    req.KycId,
	}
}

func (svc *AccountService) toCustomerProto(cus db.Customer) *pb.Customer {
	return &pb.Customer{
		Id:       fmt.Sprintf("%v", cus.ID),
		Email:    cus.Email,
		FullName: cus.FullName,
		Address:  cus.Address,
		KycType:  cus.KycType,
		KycId:    cus.KycID,
	}
}
