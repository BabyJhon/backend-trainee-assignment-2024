package entity

type User struct {
	Id       string `json:"id" validate:"uuid" db:"id"`
	Email    string `json:"email" validate:"required,email" db:"email"`
	Password string `json:"password" validate:"required,min=1" db:"password_hash"`
	UserType string `json:"user_type" validate:"required,oneof=client moderator" db:"user_type"`
}
