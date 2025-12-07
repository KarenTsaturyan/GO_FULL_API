package req

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	// we are gonna use Validator pckg
	// which works with Struct Tags in struct
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
