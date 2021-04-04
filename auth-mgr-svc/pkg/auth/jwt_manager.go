package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/atyagi9006/banking-app/auth-mgr-svc/db"
	"github.com/atyagi9006/banking-app/auth-mgr-svc/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

// jwtManager is a JSON web token manager
type jwtManager struct {
	secretKey        string
	accessSecretKey  string
	refreshSecretKey string
	tokenDuration    time.Duration
}

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(c *config.JWtConfig) *jwtManager {
	return &jwtManager{
		secretKey:        c.SecretKey,
		tokenDuration:    tokenDuration,
		accessSecretKey:  c.SecretKey,
		refreshSecretKey: c.RefreshSecretKey,
	}
}

// Generate generates and signs a new token for a user
func (manager *jwtManager) Generate(employee *db.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username: employee.Email,
		Role:     employee.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *jwtManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	log.Printf("calms %v", claims)
	return claims, nil
}
