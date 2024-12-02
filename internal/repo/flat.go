package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/pkg/postgres"
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
	if err!=nil{
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
