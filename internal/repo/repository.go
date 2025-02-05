package repo

import (
	"context"
	"time"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Flat interface {
	CreateFlat(ctx context.Context, flat entity.Flat, updatedAt time.Time) (int, error)
	UpdateFlat(ctx context.Context, flatId, houseId int, status string) (entity.Flat, error)
}

type House interface {
	CreateHouse(ctx context.Context, house entity.House) (int, error)
	GetApprovedFlats(ctx context.Context, houseId int) (*[]entity.Flat, error)
	GetAllFlats(ctx context.Context, houseId int) (*[]entity.Flat, error)
}

type Auth interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	GetUser(ctx context.Context, id, password string) (entity.User, error)
}

type Repository struct {
	Flat
	House
	Auth
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Flat:  NewFlatRepo(db),
		House: NewHouseRepo(db),
		Auth:  NewAuthRepo(db),
	}
}
