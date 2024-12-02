package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/repo"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "asdklijasd"
	TokenTTL   = 12 * time.Hour
	signingKey = "das123890op123eawd"
)

type AuthService struct {
	repo repo.Auth
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserType string `json:"user_type"`
}

func NewAuthService(repo repo.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) CreateUser(ctx context.Context, user entity.User) (string, error) {
	user.Password = a.generatePasswordHash(user.Password)
	return a.repo.CreateUser(ctx, user)
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *AuthService) GenerateToken(userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserType: userType,
	})
	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) Login(ctx context.Context, id, password string) (string, error) {
	var token string
	passwordHash := a.generatePasswordHash(password)
	user, err := a.repo.GetUser(ctx, id, passwordHash)
	if err != nil {
		return "", err
	}
	token, err = a.GenerateToken(user.UserType)
	if err != nil {
		return "", err
	}
	return token, nil
}
