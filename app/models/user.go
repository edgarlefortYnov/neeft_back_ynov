package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID              uint   `gorm:"primaryKey"   json:"id" `
	Username        string `validate:"required,min=3,max=32" gorm:"varchar(255)" json:"username"`
	FirstName       string `validate:"required,min=3,max=32" gorm:"varchar(255)" json:"firstName"`
	LastName        string `validate:"required,min=3,max=32" gorm:"varchar(255)" json:"lastName"`
	Email           string `validate:"required,min=3,max=32" gorm:"varchar(255)" json:"email"`
	EmailVerifiedAt bool   `                                 gorm:"boolean"      json:"emailVerifiedAt"`
	Password        string `validate:"required,min=8"        gorm:"varchar(255)" json:"password"`
	RememberToken   string `                                 gorm:"varchar(100)" json:"rememberToken"`
	BirthDate       string `                                 gorm:"varchar(255)" json:"birthDate"`
	Avatar          string `                                 gorm:"varchar(255)" json:"avatar"`
	IsBan           bool   `                                 gorm:"boolean"      json:"isBan"`
	LastLoginAt     time.Time
	IsSuperAdmin    bool `                                   gorm:"boolean"      json:"isSuperAdmin"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

/*
ValidateUser implements value validations for structures and individual fields based on tags.
*/
func ValidateUser(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value().(string),
			})
		}
	}
	return errors
}
