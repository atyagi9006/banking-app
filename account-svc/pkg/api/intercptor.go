package api

import (
	"context"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func accessibleRoles() map[string][]string {
	const accountServicePath = "/proto.AccountService/"
	return map[string][]string{
		accountServicePath + "CreateBankEmployee": {"admin"},
		accountServicePath + "DeleteEmployee":     {"admin"},
		accountServicePath + "GetEmployee":        {"admin", "staff"},

		//customer api
		/* accountServicePath + "CreateCustomer": {"staff"},
		accountServicePath + "UpdateCustomer": {"staff"},
		accountServicePath + "DeleteCustomer": {"staff"},
		accountServicePath + "GetCustomer":    {"staff"}, */

		//account api
		/* accountServicePath + "CreateAccount":  {"staff"},
		accountServicePath + "LinkOwner":      {"staff"},
		accountServicePath + "GetAccount":     {"staff"},
		accountServicePath + "TransferAmount": {"staff"},
		accountServicePath + "PrintStatement": {"staff"}, */
	}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AccountService) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary new interceptor: ", info.FullMethod)

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AccountService) authorize(ctx context.Context, method string) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok { //this bypass the grpc internal calls
		clientLogin := strings.Join(md["mode"], "")
		if clientLogin == "internal" {
			return nil
		}
	}
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	res2 := strings.Split(accessToken, " ")
	claims, err := interceptor.jwtManager.Verify(res2[1])
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}
	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
