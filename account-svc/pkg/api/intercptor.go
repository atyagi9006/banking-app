package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"

	authmgrPB "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
	"github.com/atyagi9006/opa-authz/opa"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//var (
// headerForwardedMethod sets an HTTP request method in gRPC gateway
//headerForwardedMethod = "x-forwarded-method"

// headerForwardedRequestURI sets an HTTP request URI in gRPC gateway
//headerForwardedRequestURI = "x-forwarded-request-uri"
//)

type Request struct {
	Path   []string `json:"path"`
	Method string   `json:"method"`
}
type Input struct {
	Request  Request           `json:"request"`
	UserRole map[string]string `json:"user_roles"`
	UserID   string            `json:"user_id"`
}

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
	log.Println("$$$$$$ inside auth ")
	// accessibleRoles, ok := interceptor.accessibleRoles[method]
	// if !ok {
	// 	// everyone can access
	// 	return nil
	// }

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	log.Println("$$$$$$MD ctx: ", md)

	clientLogin := strings.Join(md["mode"], "")
	if clientLogin == "internal" {
		return nil
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	res2 := strings.Split(accessToken, " ")

	req := authmgrPB.TokenRequest{
		Token: res2[1],
	}
	claims, err := interceptor.authmgrClient.VerifyToken(ctx, &req)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	opaData := interceptor.getOPAData(method, claims)

	allow, err := interceptor.opaClient.Authorize(context.Background(), "opa/authz", opaData)
	log.Println("$$$$$$OPA res: ", allow)
	if allow {
		return nil
	}

	// for _, role := range accessibleRoles {
	// 	if role == claims.Role {
	// 		return nil
	// 	}
	// }
	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}

func (interceptor *AccountService) getOPAData(method string, claims *authmgrPB.VerifyTokenResponse) *bytes.Buffer {
	data := Request{
		Path: splitURI(method),
	}

	// if val, ok := md[headerForwardedMethod]; ok && len(val) > 0 {
	// 	data.Method = val[0]
	// }

	// if val, ok := md[headerForwardedRequestURI]; ok && len(val) > 0 {
	// 	data.Path = splitURI(val[0])
	// }
	roles := make(map[string]string)
	roles[claims.Email] = claims.Role

	body := opa.Input{
		Input: Input{
			Request:  data,
			UserRole: roles,
			UserID:   claims.Email,
		},
	}

	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(body); err != nil {
		log.Fatal(err)
	}

	log.Println("$$$$$$OPA DAta: ", b.String())
	log.Println("$$$$$$ exit auth ")
	return &b
}

func splitURI(s string) []string {
	return strings.Split(strings.Trim(s, "/"), "/")
}
