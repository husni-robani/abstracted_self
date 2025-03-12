package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
)

// error in package
var ErrInvalidCredetials = errors.New("invalid credentials")

type AuthService struct{}

func NewAuthService() AuthService {
	return AuthService{}
}

func (service AuthService) LoginService(username, password string) (string, error) {
	if username != os.Getenv("USERNAME") || password != os.Getenv("PASSWORD") {
		logger.Info.Printf("user failed to login: invalid credentials (username: %s | password: %s)", username, password)
		return "", ErrInvalidCredetials
	}

	// issueing JWT token
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer: "self-abstracted",
		Subject: os.Getenv("USERNAME"),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("AUTH_KEY")))
	if err != nil {
		logger.Error.Printf("failed to signing token: %v", err)
		return "", err
	}

	// save the key to user_env_creds
	os.Setenv("AUTH_TOKEN", signedToken)

	return signedToken, nil
}