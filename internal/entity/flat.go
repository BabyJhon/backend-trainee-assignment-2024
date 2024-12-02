package entity

type Flat struct {
	Id      int    `json:"id" validate:"required,gte=1"  db:"id"`
	HouseId int    `json:"house_id" validate:"required,gte=1" db:"house_id"`
	Price   int    `json:"price" validate:"required,gte=0" db:"price"`
	Rooms   int    `json:"rooms" validate:"required,gte=0" db:"rooms"`
	Status  string `json:"status" validate:"omitempty,oneof=created approved declined 'on moderation'" db:"status"`
}
