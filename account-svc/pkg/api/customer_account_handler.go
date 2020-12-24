package api

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/atyagi9006/banking-app/account-svc/db"
	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var accountType = map[string]string{
	"savings": "savings",
	"salary":  "salary",
	"loan":    "loan",
	"current": "current",
}

var currency = map[string]string{
	"USD": "USD",
	"EUR": "EUR",
	"CAD": "CAD",
	"INR": "INR",
}

const (
	errInvalidAccountType = "Invalid account type"
	errInvalidCurrency    = "Invalid currency"
	errEmptyOwner         = "Owner can't be empty"
	errOwnerAlreadyExists = "Can't Link Owner already Exists"
	errAccountNotFound    = "Account not found"
)

func (svc *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.Account, error) {

	if _, ok := accountType[req.Type]; !ok {
		return nil, status.Error(codes.InvalidArgument, errInvalidAccountType)
	}

	if _, ok := currency[req.Currency]; !ok {
		return nil, status.Error(codes.InvalidArgument, errInvalidCurrency)
	}

	dbParam := svc.fromCreateAccountProto(req)

	if req.Owner != "" { //check customer exists
		getCusReq := pb.GetCustomerRequest{
			Id: req.Owner,
		}
		_, err := svc.GetCustomer(ctx, &getCusReq)
		if err != nil {
			log.Println("Error while getting owner. Error ", err)
			return nil, err
		}

		owner, err := strconv.ParseInt(req.Owner, 10, 64)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
		}
		dbParam.Owner = owner
	}

	if req.Balance == "" {
		dbParam.Balance = 0
	} else {
		bal, err := strconv.ParseInt(req.Balance, 10, 64)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
		}
		dbParam.Balance = bal
	}

	res, err := svc.store.CreateAccount(ctx, dbParam)
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}

	acc := svc.toAccountProto(res)
	return acc, nil
}

func (svc *AccountService) LinkOwner(ctx context.Context, req *pb.LinkOwnerRequest) (*pb.Account, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	if req.Owner == "" {
		return nil, status.Error(codes.InvalidArgument, errEmptyOwner)
	}

	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	owner, err := strconv.ParseInt(req.Owner, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	existingAcc, err := svc.store.GetAccount(ctx, id)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new Account: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}

	if existingAcc.Owner != 0 {
		return nil, status.Error(codes.AlreadyExists, errOwnerAlreadyExists)
	}
	args := db.LinkOwnerParams{
		ID:    id,
		Owner: owner,
	}
	res, err := svc.store.LinkOwner(ctx, args)
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}
	acc := svc.toAccountProto(res)
	return acc, nil
}

func (svc *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	var res *pb.Account
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}
	acc, err := svc.store.GetAccount(ctx, id)
	if err != nil {

		if strings.Contains(err.Error(), errNoRows) {
			return nil, status.Error(codes.NotFound, errAccountNotFound)
		}
		log.Println("Error ", err)
		return nil, status.Error(codes.Internal, errInternal)
	}
	res = svc.toAccountProto(acc)
	return res, nil
}

func (svc *AccountService) TransferAmount(ctx context.Context, req *pb.TransferAmountRequest) (*pb.Account, error) {
	if req.Amount == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	amt, err := strconv.ParseInt(req.Amount, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	if req.FromAccountId == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	fromAccID, err := strconv.ParseInt(req.FromAccountId, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	if req.ToAccountId == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	toAccID, err := strconv.ParseInt(req.ToAccountId, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}

	fromAccount, err := svc.store.GetAccount(ctx, fromAccID)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new Account: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}

	toAccount, err := svc.store.GetAccount(ctx, toAccID)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new Account: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}
	trxParam := db.TransferTxParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        amt,
	}
	_, err = svc.store.TransferTx(ctx, trxParam)
	if err != nil {
		return nil, status.Error(codes.Internal, errInternal)
	}

	fromAccountres, err := svc.store.GetAccount(ctx, fromAccID)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new Account: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}
	res := svc.toAccountProto(fromAccountres)
	return res, nil
}

func (svc *AccountService) PrintStatement(ctx context.Context, req *pb.PrintStatementRequest) (*pb.PrintStatementResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}
	accID, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, errInvalidArgument)
	}
	acc, err := svc.store.GetAccount(ctx, accID)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new Account: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}

	entries, err := svc.store.GetEntries(ctx, acc.ID)
	if err != nil {
		if !strings.Contains(err.Error(), errNoRows) {
			log.Println("Error with create new Account: ", err)
			return nil, status.Error(codes.Internal, errInternal)
		}
	}
	res := pb.PrintStatementResponse{
		Entries: svc.toEntriesProto(entries),
	}
	return &res, nil
}

func (svc *AccountService) fromCreateAccountProto(req *pb.CreateAccountRequest) db.CreateAccountParams {
	return db.CreateAccountParams{
		Currency: req.Currency,
		Type:     req.Type,
	}
}

func (svc *AccountService) toAccountProto(acc db.Account) *pb.Account {
	return &pb.Account{
		Id:       fmt.Sprintf("%v", acc.ID),
		Owner:    fmt.Sprintf("%v", acc.Owner),
		Balance:  fmt.Sprintf("%v", acc.Balance),
		Currency: acc.Currency,
		Type:     acc.Type,
	}
}

func (svc *AccountService) toEntriesProto(entries []db.Entry) []*pb.Entry {
	var res []*pb.Entry
	for _, val := range entries {
		entryRes := pb.Entry{
			Amount: int32(val.Amount),
			Date:   val.CreatedAt.String(),
		}
		res = append(res, &entryRes)
	}
	return res
}
