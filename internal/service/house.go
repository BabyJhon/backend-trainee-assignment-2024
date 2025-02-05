package service

import (
	"context"
	"errors"
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

func (s *HouseService) GetFlatsByHouse(ctx context.Context, houseId int, userStatus string) (*[]entity.Flat, error) {
	var flats *[]entity.Flat
	var err error

	if userStatus == "moderator" {
		flats, err = s.repo.GetAllFlats(ctx, houseId)
		if err != nil {
			return nil, err
		}
	} else if userStatus == "client" {
		flats, err = s.repo.GetApprovedFlats(ctx, houseId)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("wrong user type")
	}

	return flats, nil
}