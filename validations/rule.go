package validations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func FirstName() []validation.Rule {
	var firstName = []validation.Rule{
		validation.Required,
		validation.Length(2, 30),
	}
	return firstName
}

// func UniqueEmail() validation.Rule {
// 	return validation.By(func(value interface{}) error {
// 		email, ok := value.(*string)
// 		if !ok {
// 			return fmt.Errorf("expected a string, got %T", value)
// 		}
// 		emailString := *email

// 		fmt.Println("Email received for validation:", emailString)

// 		return fmt.Errorf("expected a string, got %T", value)
// 	})
// }
