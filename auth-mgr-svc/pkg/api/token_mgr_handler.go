package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	pb "github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/proto"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errInvaildUserNamePassword = "Incorrect username or password"
	errInternal                = "Internal Error"
	errInvalidEmail            = "Invalid email"
	errInvalidPassword         = "Invalid Password"
	errInvalidArgument         = "Invalid Argument"
	errInvalidRole             = "Invalid Role"
	errNoRows                  = "no rows"
	errUserExists              = "User already Exists"
	errUserNotFound            = "User not found"
)

func (svc *AuthMgrService) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	if !govalidator.IsEmail(req.Email) {
		return nil, status.Error(codes.InvalidArgument, errInvalidEmail)
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, errInvalidPassword)
	}

	emp, err := svc.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Println("Error ", err)
		if strings.Contains(err.Error(), errNoRows) {
			return nil, status.Error(codes.NotFound, errUserNotFound)
		}
		return nil, status.Error(codes.Internal, errInternal)
	}

	if req.Password != emp.Password {
		return nil, status.Error(codes.InvalidArgument, errInvaildUserNamePassword)
	}

	tokenDetails, err := svc.jwtManager.CreateToken(&emp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	err = svc.authStore.CreateAuth(strconv.Itoa(int(emp.ID)), tokenDetails)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := pb.GenerateTokenResponse{
		AccessToken:  tokenDetails.AccessToken,
		RefreshToken: tokenDetails.RefreshToken,
	}
	return &res, nil
}

func (svc *AuthMgrService) VerifyToken(ctx context.Context, req *pb.TokenRequest) (*pb.VerifyTokenResponse, error) {
	claims, err := svc.jwtManager.ExtractAccessTokenMetadata(req.Token)
	if err != nil {
		log.Println("Error while verify token:", err)
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	_, err = svc.authStore.FetchAuth(claims.TokenUuid)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}
	res := pb.VerifyTokenResponse{
		Email: claims.Username,
		Role:  claims.Role,
	}
	return &res, nil
}

func (svc *AuthMgrService) ExpireToken(ctx context.Context, req *pb.TokenRequest) (*pb.EmptyMessageResponse, error) {
	//If metadata is passed and the tokens valid, delete them from the redis store
	metadata, _ := svc.jwtManager.ExtractAccessTokenMetadata(req.Token)
	if metadata != nil {
		deleteErr := svc.authStore.DeleteTokens(metadata)
		if deleteErr != nil {
			log.Printf("Error while delete %v", deleteErr)

			return nil, status.Errorf(codes.InvalidArgument, "bad logout request")
		}
	}
	//return success result
	return &pb.EmptyMessageResponse{}, nil
}

func (svc *AuthMgrService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.GenerateTokenResponse, error) {
	var res pb.GenerateTokenResponse
	//verify the token
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		log.Printf("error while verify refresh token %v ", err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		log.Printf("Error invalid refresh token claims : %v ", err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {

			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}
		userId, roleOk := claims["user_id"].(string)
		if roleOk == false {
			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}
		//Delete the previous Refresh Token
		delErr := svc.authStore.DeleteRefresh(refreshUuid)
		if delErr != nil { //if any goes wrong
			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}

		orUserId, err := strconv.Atoi(userId)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
		}

		//need to get user from db
		user, err := svc.store.GetUser(ctx, int64(orUserId))
		if err != nil {
			log.Println("Error ", err)
			if strings.Contains(err.Error(), errNoRows) {
				return nil, status.Error(codes.NotFound, errUserNotFound)
			}
			return nil, status.Error(codes.Internal, errInternal)
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := svc.jwtManager.CreateToken(&user)
		if createErr != nil {
			return nil, status.Errorf(codes.Internal, createErr.Error())
		}
		//save the tokens metadata to redis
		saveErr := svc.authStore.CreateAuth(userId, ts)
		if saveErr != nil {
			return nil, status.Errorf(codes.Internal, saveErr.Error())
		}

		res = pb.GenerateTokenResponse{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}
	} else {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}
	return &res, nil

}
