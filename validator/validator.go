package validator

import (
	"log"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gofiber/fiber/v3"
)

// simple struct validator for testing
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

func (*structValidator) Engine() any {
	return ""
}

// ValidateStruct ...
func (v *structValidator) ValidateStruct(out any) error {
	if err := v.validator.Struct(out); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			if len(validationErrors) > 0 {
				return fiber.NewError(fiber.StatusBadRequest, validationErrors[0].Translate(v.trans))
			}
		}
	}

	return nil
}
