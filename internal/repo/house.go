package repo

import (
	"context"
	"fmt"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HouseRepo struct {
	db *pgxpool.Pool
}

func NewHouseRepo(db *pgxpool.Pool) *HouseRepo {
	return &HouseRepo{
		db: db,
	}
}

func (r *HouseRepo) CreateHouse(ctx context.Context, house entity.House) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (address, year, developer, created_at, updated_at) values ($1, $2, $3, $4, $5) RETURNING id", postgres.HousesTable)

	row := r.db.QueryRow(ctx, query, house.Address, house.Year, house.Developer, house.CreatedAt, house.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *HouseRepo) GetAllFlats(ctx context.Context, houseId int) (*[]entity.Flat, error) {
	var flats []entity.Flat
	query := fmt.Sprintf("SELECT * FROM %s WHERE house_id = $1", postgres.FlatsTable)

	rows, err := r.db.Query(ctx, query, houseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var flat entity.Flat
		if err := rows.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status); err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}

	return &flats, nil
}

func (r *HouseRepo) GetApprovedFlats(ctx context.Context, houseId int) (*[]entity.Flat, error) {
	var flats []entity.Flat
	query := fmt.Sprintf("SELECT * FROM %s WHERE house_id = $1 AND status = 'approved'", postgres.FlatsTable)

	rows, err := r.db.Query(ctx, query, houseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var flat entity.Flat
		if err := rows.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status); err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}

	return &flats, nil
}
