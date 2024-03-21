package validate

import (
	"log"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gofiber/fiber/v3"
)

var (
	v     *validator.Validate
	trans ut.Translator
)

func init() {
	v = validator.New()
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	err := translations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		log.Println(err.Error())
	}
}

// Struct validate params struct
func Struct(params interface{}) error {
	if err := v.Struct(params); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			if len(validationErrors) > 0 {
				return fiber.NewError(fiber.StatusBadRequest, validationErrors[0].Translate(trans))
			}
		}
	}

	return nil
}

// Variable Single parameter
func Variable(fieldName string, val interface{}, tag string) error {
	if err := v.Var(val, strings.ReplaceAll(tag, " ", "")); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			if len(validationErrors) > 0 {
				return fiber.NewError(fiber.StatusBadRequest, fieldName+validationErrors[0].Translate(trans))
			}
		}
	}

	return nil
}
