package api

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/atyagi9006/banking-app/account-svc/auth"
	"github.com/atyagi9006/banking-app/account-svc/db"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:P@ssw0rd@localhost:5432/my_bank?sslmode=disable"
)

type AccountService struct {
	store *db.Store

	jwtManager      *auth.JWTManager
	accessibleRoles map[string][]string
}

func NewAccountService() *AccountService {
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(testDB)
	jwtMangager := auth.NewJWTManager(auth.SecretKey, auth.TokenDuration)
	accountService := AccountService{
		store:           store,
		jwtManager:      jwtMangager,
		accessibleRoles: accessibleRoles(),
	}
	return &accountService
}

func accessibleRoles() map[string][]string {
	const accountServicePath = "/proto.AccountService/"

	return map[string][]string{
		accountServicePath + "CreateBankEmployee": {"admin"},
		accountServicePath + "DeleteEmployee":     {"admin"},
		accountServicePath + "GetEmployee":        {"admin", "staff"},
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
