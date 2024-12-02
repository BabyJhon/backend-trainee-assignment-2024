package service

import (
	"context"
	"time"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/repo"
)

type HouseService struct {
	repo repo.House
}

func NewHouseService(repo repo.House) *HouseService {
	return &HouseService{
		repo: repo,
	}
}


func (s *HouseService) CreateHouse(ctx context.Context, house entity.House) (entity.House, error) {
	house.CreatedAt = time.Now()
	house.UpdatedAt = time.Now()
	id, err := s.repo.CreateHouse(ctx, house)
	if err != nil {
		return entity.House{}, err
	}
	house.Id = id
	return house, nil
}
