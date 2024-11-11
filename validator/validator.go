package validator

import (
	"log"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gofiber/fiber/v3"
)

// Custom validator
type structValidator struct {
	trans     ut.Translator
	validator *validator.Validate
}

// New creates a new validator
func New() *structValidator {
	validator := validator.New()
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	err := translations.RegisterDefaultTranslations(validator, trans)
	if err != nil {
		log.Fatal(err)
	}

	return &structValidator{
		trans:     trans,
		validator: validator,
	}
}

// Validate ...
func (v *structValidator) Validate(out any) (err error) {
	err = v.validator.Struct(out)
	if err == nil {
		return
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok && len(validationErrors) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, validationErrors[0].Translate(v.trans))
	}

	return
}
