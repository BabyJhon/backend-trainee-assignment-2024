package entity

import (
	"time"
)

type House struct {
	Id        int       `json:"id" db:"id"`
	Address   string    `json:"address" validate:"required,min=1" db:"address"`
	Year      int       `json:"year" validate:"required,gte=0" db:"year"`
	Developer string    `json:"developer" validate:"omitempty,min=1" db:"developer"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
