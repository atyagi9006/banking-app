package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/atyagi9006/banking-app/auth-mgr-svc/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type TokenProvider interface {
	CreateToken(user *db.User) (*TokenDetails, error)
	ExtractAccessTokenMetadata(AccessToken string) (*AccessDetails, error)
}

//Token implements the TokenInterface
var _ TokenProvider = &jwtManager{}

func (t *jwtManager) CreateToken(user *db.User) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix() //expires after 30 min
	td.TokenUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = user.ID
	atClaims["user_email"] = user.Email
	atClaims["role"] = user.Role
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(t.accessSecretKey))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + fmt.Sprintf("%d", user.ID)

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = user.ID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(t.refreshSecretKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func extract(token *jwt.Token) (*AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("something went wrong")
	}
	accessUUID, ok := claims["access_uuid"].(string)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	userID, userOk := claims["user_id"].(string)
	if !userOk {
		return nil, errors.New("unauthorized")
	}
	userName, userNameOk := claims["user_id"].(string)
	if !userNameOk {
		return nil, errors.New("unauthorized")
	}

	role, roleOk := claims["role"].(string)
	if !roleOk {
		return nil, errors.New("unauthorized")
	}

	ad := AccessDetails{
		TokenUuid: accessUUID,
		UserId:    userID,
		Username:  userName,
		Role:      role,
	}
	return &ad, nil
}

func (t *jwtManager) ExtractAccessTokenMetadata(AccessToken string) (*AccessDetails, error) {
	token, err := verifyToken(AccessToken)
	if err != nil {
		return nil, err
	}
	acc, err := extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
