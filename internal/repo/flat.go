package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/pkg/postgres"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FlatRepo struct {
	db *pgxpool.Pool
}

func NewFlatRepo(db *pgxpool.Pool) *FlatRepo {
	return &FlatRepo{
		db: db,
	}
}

func (r *FlatRepo) CreateFlat(ctx context.Context, flat entity.Flat, updatedAt time.Time) (int, error) {
	var id int
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (id, house_id, price, rooms, status) values ($1, $2, $3, $4, $5) RETURNING id", postgres.FlatsTable)

	row := tx.QueryRow(ctx, query, flat.Id, flat.HouseId, flat.Price, flat.Rooms, flat.Status)
	if err := row.Scan(&id); err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	updateHouse := fmt.Sprintf("UPDATE %s SET updated_at = $1 WHERE id = $2", postgres.HousesTable)
	_, err = tx.Exec(ctx, updateHouse, updatedAt, id)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	return id, tx.Commit(ctx)
}

func (r *FlatRepo) UpdateFlat(ctx context.Context, flatId, houseId int, status string) (entity.Flat, error) {
	var flat entity.Flat

	onModeration, err := r.isFlatOnModeration(ctx, flatId, houseId)
	if err != nil {
		return entity.Flat{}, err
	}
	if onModeration {
		return entity.Flat{}, ErrFlatOnModeration
	}
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return entity.Flat{}, err
	}

	updateFlat := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2 AND house_id = $3 RETURNING id, house_id, price, rooms, status", postgres.FlatsTable)
	row := tx.QueryRow(ctx, updateFlat, status, flatId, houseId)
	if err := row.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status); err != nil {
		//fmt.Println(flat)
		tx.Rollback(ctx)
		return entity.Flat{}, err
	}

	updatedAt := time.Now()
	updateHouse := fmt.Sprintf("UPDATE %s SET updated_at = $1 WHERE id = $2", postgres.HousesTable)
	_, err = tx.Exec(ctx, updateHouse, updatedAt, houseId)
	if err != nil {
		tx.Rollback(ctx)
		return entity.Flat{}, err
	}

	return flat, tx.Commit(ctx)
}

func (r *FlatRepo) isFlatOnModeration(ctx context.Context, flatId, houseId int) (bool, error) {
	var flat entity.Flat
	query := fmt.Sprintf("SELECT status FROM %s WHERE id = $1 AND house_id = $2", postgres.FlatsTable)
	row := r.db.QueryRow(ctx, query, flatId, houseId)
	if err := row.Scan(&flat.Status); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, ErrFlatNotFound
		}
		return false, err
	}
	if flat.Status == "on moderation" {
		return true, nil
	}
	return false, nil
}
