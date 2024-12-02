package service

import (
	"context"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/repo"
)

type Flat interface {
	CreateFlat(ctx context.Context, flat entity.Flat) (entity.Flat, error) 
}

type House interface {
	CreateHouse(ctx context.Context, house entity.House) (entity.House, error)
}

type Auth interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	generatePasswordHash(password string) string 
	GenerateToken(userType string) (string, error)
	Login(ctx context.Context, id, password string) (string, error) 
}

type Service struct {
	Flat
	House
	Auth
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Flat: NewFlatService(repos),
		House: NewHouseService(repos),
		Auth: NewAuthService(repos),
	}
}
