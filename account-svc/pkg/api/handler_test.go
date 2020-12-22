package api

import (
	"context"
	"testing"

	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateEmployee(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"create Employee with no data or empty email ": func() {
			svc := NewAccountService()
			req := pb.CreateEmployeeRequest{}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidEmail)
			assert.Empty(t, resp)
		},
		"create Employee with invalid email": func() {
			svc := NewAccountService()
			req := pb.CreateEmployeeRequest{Email: "test@t"}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidEmail)
			assert.Empty(t, resp)
		},
		"create Employee with invalid Password": func() {
			svc := NewAccountService()
			req := pb.CreateEmployeeRequest{Email: "test@t.com", Password: ""}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidPassword)
			assert.Empty(t, resp)
		},
		"create Employee with invalid/empty FullName": func() {
			svc := NewAccountService()
			req := pb.CreateEmployeeRequest{Email: "test@t.com",
				Password: "tes@password",
				FullName: "",
			}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			//assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errEmptyFullName)
			assert.Empty(t, resp)
		},
		"create Employee with Empty Role": func() {
			svc := NewAccountService()
			req := pb.CreateEmployeeRequest{Email: "test@t.com",
				Password: "tes@password",
				FullName: "TestName",
			}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			//assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidRole)
			assert.Empty(t, resp)
		},
		"create Employee with Invalid Role": func() {
			svc := NewAccountService()
			req := pb.CreateEmployeeRequest{Email: "test@t.com",
				Password: "tes@password",
				FullName: "TestName",
				Role:     "Test",
			}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			//assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidRole)
			assert.Empty(t, resp)
		},
		"create Employee Success": func() {
			svc := NewAccountService()
			CreateRamdomUser(t, svc)
		},
		"create already existing Employee": func() {
			svc := NewAccountService()
			emp := CreateRamdomUser(t, svc)
			req := pb.CreateEmployeeRequest{Email: emp.Email,
				Password: "tes@password",
				FullName: emp.FullName,
				Role:     randomRole(),
			}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			assert.Nil(t, resp)
			assert.Equal(t, codes.AlreadyExists, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errEmployeeExists)
		},
	}
	t.Log("Create user tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func CreateRamdomUser(t *testing.T, svc *AccountService) *pb.Employee {
	req := pb.CreateEmployeeRequest{Email: randomEmail(),
		Password: randomPassword(),
		FullName: randomFullName(),
		Role:     randomRole(),
	}
	ctx := context.Background()
	resp, err := svc.CreateBankEmployee(ctx, &req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id)
	assert.Equal(t, req.Email, resp.Email)
	assert.Equal(t, req.FullName, resp.FullName)
	assert.Equal(t, req.Role, resp.Role)

	return resp
}

func TearDownRandomUser(t *testing.T, svc *AccountService, id string) {

}
