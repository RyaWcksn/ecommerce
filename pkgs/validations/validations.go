package validations

import (
	"fmt"

	"github.com/go-playground/validator"
)

func Validate(s any) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("[ValidateCreateBadgeRequest] Invalid request in field [%+v], tag [%+v], value [%+v]", err.StructNamespace(), err.Tag(), err.Param())
		}
	}
	return nil
}
