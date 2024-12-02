package middleware

import (
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func IsCreateHouseInputValid(house entity.House) (bool, error) {
	validate = validator.New()
	if err := validate.Struct(house); err != nil {
		return false, err
	}
	return true, nil
}

func IsDummyLoginInputValid(userType string) (bool, error) {
	validate := validator.New()
	if err := validate.Var(userType, "oneof=client moderator"); err != nil {
		return false, err
	}
	return true, nil
}

func IsRegisterInputValid(user entity.User) (bool, error) {
	validate = validator.New()
	if err := validate.Struct(user); err != nil {
		return false, err
	}
	return true, nil
}

func IsCreateFlatInputValid(flat entity.Flat) (bool, error) {
	validate = validator.New()
	if err := validate.Struct(flat); err != nil {
		return false, err
	}
	return true, nil
}

func IsLoginInputValid(id, password string) (bool, error) {
	validate = validator.New()
	if err := validate.Var(id, `"required,uuid"`); err != nil {
		return false, err
	}
	if err := validate.Var(password, `"required,min=1"`); err != nil {
		return false, err
	}
	return true, nil
}