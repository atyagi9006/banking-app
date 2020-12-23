package api

import (
	"context"
	"testing"

	pb "github.com/atyagi9006/banking-app/account-svc/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateRamdomEmployee(t *testing.T, svc *AccountService) *pb.Employee {
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
	ctx := context.Background()
	req := pb.EmployeeIdRequest{Id: id}
	_, err := svc.DeleteEmployee(ctx, &req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateEmployee(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"create Employee with no data or empty email ": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateEmployeeRequest{}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidEmail)
			assert.Empty(t, resp)
		},
		"create Employee with invalid email": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateEmployeeRequest{Email: "test@t"}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidEmail)
			assert.Empty(t, resp)
		},
		"create Employee with invalid Password": func() {
			//setup
			svc := NewAccountService()

			//test
			req := pb.CreateEmployeeRequest{Email: "test@t.com", Password: ""}
			ctx := context.Background()
			resp, err := svc.CreateBankEmployee(ctx, &req)
			assert.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Convert(err).Code())
			assert.Contains(t, err.Error(), errInvalidPassword)
			assert.Empty(t, resp)
		},
		"create Employee with invalid/empty FullName": func() {
			//setup
			svc := NewAccountService()

			//test
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
			assert.Contains(t, err.Error(), errEmptyFullName)
			assert.Empty(t, resp)
		},
		"create Employee with Empty Role": func() {
			//setup
			svc := NewAccountService()

			//test
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
			assert.Contains(t, err.Error(), errInvalidRole)
			assert.Empty(t, resp)
		},
		"create Employee with Invalid Role": func() {
			//setup
			svc := NewAccountService()

			//test
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
			assert.Contains(t, err.Error(), errInvalidRole)
			assert.Empty(t, resp)
		},
		"create Employee Success": func() {
			//setup
			svc := NewAccountService()

			//test
			emp := CreateRamdomEmployee(t, svc)

			//tear down
			TearDownRandomUser(t, svc, emp.Id)
		},
		"create already existing Employee": func() {
			//setup
			svc := NewAccountService()
			emp := CreateRamdomEmployee(t, svc)

			//test
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

			//teardown
			TearDownRandomUser(t, svc, emp.Id)
		},
	}
	t.Log("Create Employee tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestGetEmployee(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"Get Employee with invalid argument": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetEmployee(ctx, &pb.GetEmployeeRequest{Id: ""})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"Get Employee with invalid email": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetEmployee(ctx, &pb.GetEmployeeRequest{Email: "abc@c"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidEmail)
			assert.Empty(t, resp)
		},
		"Get Non Existing Employee With email": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetEmployee(ctx, &pb.GetEmployeeRequest{Email: "abc@c.com"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.NotFound, e.Code())
			}

			assert.Contains(t, err.Error(), errEmployeeNotFound)
			assert.Empty(t, resp)
		},
		"Get Non Existing Employee With id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetEmployee(ctx, &pb.GetEmployeeRequest{Id: "124"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.NotFound, e.Code())
			}
			assert.Contains(t, err.Error(), errEmployeeNotFound)
			assert.Empty(t, resp)
		},
		"Get Employee with invalid id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.GetEmployee(ctx, &pb.GetEmployeeRequest{Id: "t-123"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"Get Employee With id success ": func() {
			//setup
			svc := NewAccountService()
			emp := CreateRamdomEmployee(t, svc)

			//test
			ctx := context.Background()
			resp, err := svc.GetEmployee(ctx, &pb.GetEmployeeRequest{Id: emp.Id})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Id)
			assert.Equal(t, emp.Id, resp.Id)
			assert.Equal(t, emp.Email, resp.Email)
			assert.Equal(t, emp.FullName, resp.FullName)
			assert.Equal(t, emp.Role, resp.Role)

			//tear Down
			TearDownRandomUser(t, svc, emp.Id)
		},
		"Get Employee With email success ": func() {
			//setup
			svc := NewAccountService()
			emp := CreateRamdomEmployee(t, svc)

			//test
			ctx := context.Background()
			resp, err := svc.GetEmployee(ctx, &pb.GetEmployeeRequest{Email: emp.Email})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotEmpty(t, resp.Id)
			assert.Equal(t, emp.Id, resp.Id)
			assert.Equal(t, emp.Email, resp.Email)
			assert.Equal(t, emp.FullName, resp.FullName)
			assert.Equal(t, emp.Role, resp.Role)

			//tear Down
			TearDownRandomUser(t, svc, emp.Id)
		},
	}
	t.Log("Get Employee tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	type testBuilder func()
	testCases := map[string]testBuilder{
		"Delete Employee with empty id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.DeleteEmployee(ctx, &pb.EmployeeIdRequest{Id: ""})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
			assert.Nil(t, resp)
		},
		"Delete Employee with invalid id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.DeleteEmployee(ctx, &pb.EmployeeIdRequest{Id: "t-123"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.InvalidArgument, e.Code())
			}
			assert.Contains(t, err.Error(), errInvalidArgument)
			assert.Empty(t, resp)
		},
		"Delete Non Existing Employee With id": func() {
			//setup
			svc := NewAccountService()

			//test
			ctx := context.Background()
			resp, err := svc.DeleteEmployee(ctx, &pb.EmployeeIdRequest{Id: "124"})
			assert.Error(t, err)
			if e, ok := status.FromError(err); ok {
				assert.Equal(t, codes.NotFound, e.Code())
			}
			assert.Contains(t, err.Error(), errEmployeeNotFound)
			assert.Empty(t, resp)
		},
		"Delete Employee With id success ": func() {
			//setup
			svc := NewAccountService()
			emp := CreateRamdomEmployee(t, svc)

			//test
			ctx := context.Background()
			resp, err := svc.DeleteEmployee(ctx, &pb.EmployeeIdRequest{Id: emp.Id})
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		},
	}

	t.Log("Get Employee tests")
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			testCase()
		})
	}
}
