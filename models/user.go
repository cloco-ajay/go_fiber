package models

import (
	"sales-api/validations"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID              uint       `json:"id" gorm:"primaryKey; autoIncrement"`
	Name            string     `json:"name"`
	Email           *string    `json:"email" gorm:"unique"` //allowing null values
	Password        string     `json:"password"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

func (u User) Validate(create bool) error {
	var validationFields []*validation.FieldRules

	validationFields = append(validationFields,
		validation.Field(&u.Name, validations.FirstName()...),
	)
	if create {
		validationFields = append(validationFields,
			validation.Field(&u.Email, validation.Required, is.EmailFormat),
			validation.Field(&u.Password, validation.Required, validation.Length(6, 20)),
		)
	}
	return validation.ValidateStruct(&u, validationFields...)

}
