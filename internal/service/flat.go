package service

import (
	"context"
	"time"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/repo"
)

type FlatService struct {
	repo repo.Flat
}

func NewFlatService(repo repo.Flat) *FlatService {
	return &FlatService{
		repo: repo,
	}
}

func (s *FlatService) CreateFlat(ctx context.Context, flat entity.Flat) (entity.Flat, error) {
	flat.Status = "created"
	id, err := s.repo.CreateFlat(ctx, flat, time.Now())
	if err != nil {
		return entity.Flat{}, err
	}
	flat.Id = id
	return flat, nil
}

func (s *FlatService) UpdateFlat(ctx context.Context, flatId, houseId int, status string) (entity.Flat, error) {

	flat, err := s.repo.UpdateFlat(ctx, flatId, houseId, status)
	if err != nil {
		return entity.Flat{}, err
	}

	return flat, nil
}
