package api

import (
	"context"
	"strconv"
	"testing"

	db "github.com/atyagi9006/banking-app/account-svc/db"
	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateRamdomAccount(t *testing.T, svc *AccountService) *pb.Account {
	req := pb.CreateAccountRequest{
		Balance:  db.RandomMoney(),
		Currency: db.RandomCurrency(),
		Type:     db.RandomAccountType(),
	}
	ctx := context.Background()
	resp, err := svc.CreateAccount(ctx, &req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id)
	assert.Equal(t, req.Balance, resp.Balance)
	assert.Equal(t, req.Currency, resp.Currency)
	assert.Equal(t, req.Type, resp.Type)
	return resp
}

func TearDownRandomAccount(t *testing.T, svc *AccountService, id string) {
	ctx := context.Background()
	newID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	err = svc.store.DeleteAccount(ctx, newID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateAccount(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"create Account with no data or empty AccountType ": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateAccountRequest{}
			ctx := context.Background()
			resp, err := svc.CreateAccount(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidAccountType)
			assert.Empty(t, resp)
		},
		"create Account with invalid AccountType": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateAccountRequest{Type: "test@t"}
			ctx := context.Background()
			resp, err := svc.CreateAccount(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidAccountType)
			assert.Empty(t, resp)
		},
		"create Account with invalid/empty Currency": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateAccountRequest{
				Type:     db.RandomAccountType(),
				Currency: ""}
			ctx := context.Background()
			resp, err := svc.CreateAccount(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidCurrency)
			assert.Empty(t, resp)
		},
		"create Account Success": func() {
			//setup
			svc := NewAccountService()

			//test
			acc := CreateRamdomAccount(t, svc)

			//tear down
			TearDownRandomAccount(t, svc, acc.Id)
		},
	}
	t.Log("Create Account tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestGetAccount(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"Get Account with invalid argument": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetAccount(ctx, &pb.GetAccountRequest{Id: ""})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"Get Non Existing Account With id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetAccount(ctx, &pb.GetAccountRequest{Id: "124"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.NotFound, e.Code())
			}
			assert.Contains(t, err.Error(), errAccountNotFound)
			assert.Empty(t, resp)
		},
		"Get Account with invalid id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetAccount(ctx, &pb.GetAccountRequest{Id: "t-123"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"Get Account With id success ": func() {
			//setup
			svc := NewAccountService()
			acc := CreateRamdomAccount(t, svc)

			//test
			ctx := context.Background()
			resp, err := svc.GetAccount(ctx, &pb.GetAccountRequest{Id: acc.Id})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Id)
			assert.Equal(t, acc.Id, resp.Id)
			assert.Equal(t, acc.Type, resp.Type)
			assert.Equal(t, acc.Currency, resp.Currency)
			assert.Equal(t, acc.Balance, resp.Balance)

			//tear Down
			TearDownRandomAccount(t, svc, acc.Id)
		},
	}
	t.Log("Get Account tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestLinkOwner(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"link Owner with no data  ": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.LinkOwnerRequest{}
			ctx := context.Background()
			resp, err := svc.LinkOwner(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"link Owner by invalid/empty  account id": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.LinkOwnerRequest{
				Id: "",
			}
			ctx := context.Background()
			resp, err := svc.LinkOwner(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"link Owner with Empty owner": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.LinkOwnerRequest{
				Id: "test@t.com",
			}
			ctx := context.Background()
			resp, err := svc.LinkOwner(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errEmptyOwner)
			assert.Empty(t, resp)
		},
		"link Owner Success": func() {
			//setup
			svc := NewAccountService()
			acc := CreateRamdomAccount(t, svc)
			customer := CreateRamdomCustomer(t, svc)

			//test
			req := pb.LinkOwnerRequest{
				Id:    acc.Id,
				Owner: customer.Id,
			}
			ctx := context.Background()
			resp, err := svc.LinkOwner(ctx, &req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Id)
			assert.NotEmpty(t, resp.Id)
			assert.Equal(t, acc.Id, resp.Id)
			assert.Equal(t, acc.Type, resp.Type)
			assert.Equal(t, acc.Currency, resp.Currency)
			assert.Equal(t, acc.Balance, resp.Balance)
			assert.Equal(t, customer.Id, resp.Owner)

			//teardown
			TearDownRandomAccount(t, svc, acc.Id)
			TearDownRandomCustomer(t, svc, customer.Id)
		},
	}
	t.Log("Update Account tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}
