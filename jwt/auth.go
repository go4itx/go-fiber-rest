package jwt

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
)

// New ...
func New(secret string) fiber.Handler {
	return keyauth.New(keyauth.Config{
		Validator: func(ctx fiber.Ctx, tokenString string) (bool, error) {
			claims, err := ParseToken(tokenString, secret)
			if err != nil {
				return false, err
			}

			ctx.Locals("user", claims)
			return true, nil
		},
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			if err != nil {
				err = fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}

			return err
		},
	})
}
