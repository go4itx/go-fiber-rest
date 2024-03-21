package validate

import (
	"github.com/gofiber/fiber/v3"
)

// Bind: Parse data and validate struct
func Bind(ctx fiber.Ctx, data interface{}, query ...bool) (err error) {
	if (len(query) > 0 && query[0]) || ctx.Request().Header.IsGet() {
		err = ctx.Bind().Query(data)
	} else {
		err = ctx.Bind().Body(data)
	}

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = Struct(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return
}
