package api

import (
	"context"
	"testing"

	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateRamdomCustomer(t *testing.T, svc *AccountService) *pb.Customer {
	req := pb.CreateCustomerRequest{
		Email:    randomEmail(),
		FullName: randomFullName(),
		Address:  randomAddress(),
		KycType:  randomKycType(),
		KycId:    randomKycID(),
	}
	ctx := context.Background()
	resp, err := svc.CreateCustomer(ctx, &req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id)
	assert.Equal(t, req.Email, resp.Email)
	assert.Equal(t, req.FullName, resp.FullName)
	assert.Equal(t, req.Address, resp.Address)
	assert.Equal(t, req.KycType, resp.KycType)
	assert.Equal(t, req.KycId, resp.KycId)

	return resp
}

func TearDownRandomCustomer(t *testing.T, svc *AccountService, id string) {
	ctx := context.Background()
	req := pb.CustomerIdRequest{Id: id}
	_, err := svc.DeleteCustomer(ctx, &req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateCustomer(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"create Customer with no data or empty email ": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateCustomerRequest{}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidEmail)
			assert.Empty(t, resp)
		},
		"create Customer with invalid email": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateCustomerRequest{Email: "test@t"}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidEmail)
			assert.Empty(t, resp)
		},
		"create Customer with invalid/empty Address": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateCustomerRequest{
				Email:    "test@t.com",
				FullName: "test",
				Address:  ""}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errEmptyAddress)
			assert.Empty(t, resp)
		},
		"create Customer with invalid/empty FullName": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateCustomerRequest{Email: "test@t.com",
				FullName: "",
			}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errEmptyFullName)
			assert.Empty(t, resp)
		},
		"create Customer with Empty KycType": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateCustomerRequest{Email: "test@t.com",
				Address:  randomAddress(),
				FullName: "TestName",
			}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidKycType)
			assert.Empty(t, resp)
		},
		"create Customer with Invalid KycType": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateCustomerRequest{Email: "test@t.com",
				Address:  randomAddress(),
				FullName: "TestName",
				KycType:  "Test",
			}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidKycType)
			assert.Empty(t, resp)
		},
		"create Customer with Empty KycID": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateCustomerRequest{Email: "test@t.com",
				Address:  randomAddress(),
				FullName: "TestName",
				KycType:  randomKycType(),
			}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errEmptyKycID)
			assert.Empty(t, resp)
		},
		"create Customer Success": func() {
			//setup
			svc := NewAccountService()

			//test
			cus := CreateRamdomCustomer(t, svc)

			//tear down
			TearDownRandomCustomer(t, svc, cus.Id)
		},
		"create already existing Customer": func() {
			//setup
			svc := NewAccountService()
			cus := CreateRamdomCustomer(t, svc)

			//test
			req := pb.CreateCustomerRequest{Email: cus.Email,
				Address:  randomAddress(),
				FullName: cus.FullName,
				KycType:  randomKycType(),
				KycId:    randomKycID(),
			}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			assert.Nil(t, resp)
			assert.Equal(t, codes.AlreadyExists, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errCustomerExists)

			//teardown
			TearDownRandomCustomer(t, svc, cus.Id)
		},
	}
	t.Log("Create Customer tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestGetCustomer(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"Get Customer with invalid argument": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Id: ""})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"Get Customer with invalid email": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Email: "abc@c"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidEmail)
			assert.Empty(t, resp)
		},
		"Get Non Existing Customer With email": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Email: "abc@c.com"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.NotFound, e.Code())
			}

			assert.Contains(t, err.Error(), errCustomerNotFound)
			assert.Empty(t, resp)
		},
		"Get Non Existing Customer With id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Id: "124"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.NotFound, e.Code())
			}
			assert.Contains(t, err.Error(), errCustomerNotFound)
			assert.Empty(t, resp)
		},
		"Get Non Existing Customer name": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Name: "test"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.NotFound, e.Code())
			}
			assert.Contains(t, err.Error(), errCustomerNotFound)
			assert.Empty(t, resp)
		},
		"Get Customer with invalid id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Id: "t-123"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"Get Customer With id success ": func() {
			//setup
			svc := NewAccountService()
			cus := CreateRamdomCustomer(t, svc)

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Id: cus.Id})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Id)
			assert.Equal(t, cus.Email, resp.Email)
			assert.Equal(t, cus.FullName, resp.FullName)
			assert.Equal(t, cus.KycType, resp.KycType)

			//tear Down
			TearDownRandomCustomer(t, svc, cus.Id)
		},
		"Get Customer With email success ": func() {
			//setup
			svc := NewAccountService()
			cus := CreateRamdomCustomer(t, svc)

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Email: cus.Email})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Id)
			assert.Equal(t, cus.Email, resp.Email)
			assert.Equal(t, cus.FullName, resp.FullName)
			assert.Equal(t, cus.KycType, resp.KycType)

			//tear Down
			TearDownRandomCustomer(t, svc, cus.Id)
		},
		"Get Customer With full Name success ": func() {
			//setup
			svc := NewAccountService()
			cus := CreateRamdomCustomer(t, svc)

			//test
			ctx := context.Background()
			resp, err := svc.GetCustomer(ctx, &pb.GetCustomerRequest{Name: cus.FullName})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Id)
			assert.Equal(t, cus.Email, resp.Email)
			assert.Equal(t, cus.FullName, resp.FullName)
			assert.Equal(t, cus.KycType, resp.KycType)
			assert.Equal(t, cus.KycId, resp.KycId)

			//tear Down
			TearDownRandomCustomer(t, svc, cus.Id)
		},
	}
	t.Log("Get Customer tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestDeleteCustomer(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"Delete Customer with empty id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.DeleteCustomer(ctx, &pb.CustomerIdRequest{Id: ""})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
			assert.Nil(t, resp)
		},
		"Delete Customer with invalid id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.DeleteCustomer(ctx, &pb.CustomerIdRequest{Id: "t-123"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"Delete Non Existing Customer With id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.DeleteCustomer(ctx, &pb.CustomerIdRequest{Id: "124"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.NotFound, e.Code())
			}
			assert.Contains(t, err.Error(), errCustomerNotFound)
			assert.Empty(t, resp)
		},
		"Delete Customer With id success ": func() {
			//setup
			svc := NewAccountService()
			cus := CreateRamdomCustomer(t, svc)

			//test
			ctx := context.Background()
			resp, err := svc.DeleteCustomer(ctx, &pb.CustomerIdRequest{Id: cus.Id})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		},
	}

	t.Log("Get Customer tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"update Customer with no data  ": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.UpdateCustomerRequest{}
			ctx := context.Background()
			resp, err := svc.UpdateCustomer(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"update Customer by invalid/empty id": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.UpdateCustomerRequest{
				Id: "",
			}
			ctx := context.Background()
			resp, err := svc.UpdateCustomer(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"update Customer with Empty KycType": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.UpdateCustomerRequest{
				Id: "testid",
			}
			ctx := context.Background()
			resp, err := svc.UpdateCustomer(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidKycType)
			assert.Empty(t, resp)
		},
		"update Customer with Invalid KycType": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.UpdateCustomerRequest{
				Id:      "test",
				KycType: "Test",
			}
			ctx := context.Background()
			resp, err := svc.UpdateCustomer(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidKycType)
			assert.Empty(t, resp)
		},
		"update Customer with Empty KycID": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateCustomerRequest{Email: "test@t.com",
				Address:  randomAddress(),
				FullName: "TestName",
				KycType:  randomKycType(),
			}
			ctx := context.Background()
			resp, err := svc.CreateCustomer(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errEmptyKycID)
			assert.Empty(t, resp)
		},
		"update already existing Customer Success": func() {
			//setup
			svc := NewAccountService()
			cus := CreateRamdomCustomer(t, svc)
			KycType := randomKycType()
			KycID := randomKycID()

			//test
			req := pb.UpdateCustomerRequest{
				Id:      cus.Id,
				KycType: KycType,
				KycId:   KycID,
			}
			ctx := context.Background()
			resp, err := svc.UpdateCustomer(ctx, &req)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Id)
			assert.Equal(t, cus.Email, resp.Email)
			assert.Equal(t, cus.FullName, resp.FullName)
			assert.Equal(t, cus.Address, resp.Address)
			assert.Equal(t, KycType, resp.KycType)
			assert.Equal(t, KycID, resp.KycId)

			//teardown
			TearDownRandomCustomer(t, svc, cus.Id)
		},
	}
	t.Log("Update Customer tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}
