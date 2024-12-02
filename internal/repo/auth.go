package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/BabyJhon/backend-trainee-assignment-2024/pkg/postgres"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (a *AuthRepo) CreateUser(ctx context.Context, user entity.User) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, user_type) values ($1, $2, $3) RETURNING id", postgres.UsersTable)

	row := a.db.QueryRow(ctx, query, user.Email, user.Password, user.UserType)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}



func (a *AuthRepo) GetUser(ctx context.Context, id, passwordHash string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND password_hash = $2", postgres.UsersTable)

	row := a.db.QueryRow(ctx, query, id, passwordHash)
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.UserType); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, err
	}

	return user, nil
}
