package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/rubemlrm/go-api-bootstrap/pkg/validations"
)

type UserCreate struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required,min=3"`
}

func (u *UserCreate) Validate(vf validations.ValidationFunc) error {
	vl, err := vf("en",
		validations.WithCustomFieldLabel("json"),
		validations.WithCustomValidationRule("is-awesome", ValidateMyVal),
		validations.WithCustomTranslation("is-awesome", "{0} must have a value!"),
	)

	if err != nil {
		return err
	}

	err = vl.CheckWithTranslations(u)
	if err != nil {
		return err
	}
	return nil
}

func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
