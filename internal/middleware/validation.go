package middleware

import (
	"github.com/BabyJhon/backend-trainee-assignment-2024/internal/entity"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func IsCreateHouseInputValid(house entity.House) (error) {
	validate = validator.New()
	return validate.Struct(house)
}

func IsDummyLoginInputValid(userType string) ( error) {
	validate := validator.New()
	return validate.Var(userType, "oneof=client moderator")
}

func IsRegisterInputValid(user entity.User) (error) {
	validate = validator.New()
	return validate.Struct(user)
}

func IsCreateFlatInputValid(flat entity.Flat) (error) {
	validate = validator.New()
	return validate.Struct(flat)
}

func IsLoginInputValid(id, password string) (error) {
	validate = validator.New()
	if err := validate.Var(id, "required,uuid"); err != nil {
		return err
	}
	if err := validate.Var(password, "required,min=1"); err != nil {
		return err
	}
	return nil
}