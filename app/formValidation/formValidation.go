package formValidation

import (
	"github.com/go-playground/validator/v10"
	"neeft_back/app/models"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// ValidateUserInformation validate user information
func ValidateUserInformation(userInformation models.User) []*ErrorResponse {
	var errors []*ErrorResponse
	var validateForm = validator.New()

	// Validate the user information
	err := validateForm.Struct(userInformation)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}

	return errors
}
