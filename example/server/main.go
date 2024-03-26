package main

import (
	"github.com/go4itx/go-fiber-rest/jwt"
	"github.com/go4itx/go-fiber-rest/response"
	"github.com/go4itx/go-fiber-rest/server"
	"github.com/gofiber/fiber/v3"
)

func main() {
	server.New(router)
}

// router ...
func router(app *fiber.App) {
	app.All("/", func(ctx fiber.Ctx) error {
		return response.New(ctx).JSON("Hello, World!")
	})

	app.Post("/login", func(ctx fiber.Ctx) error {
		var user User
		if err := ctx.Bind().Body(&user); err != nil {
			return err
		}

		exp := jwt.GetExpTime(7) // 设置token过期时间为7天
		return response.New(ctx).JSON(jwt.CreateToken(user.Name, exp, secret))
	})

	// 中间件，开启验证token，下面的请求需要Authorization
	app.Use(jwt.New(secret))
	// 如：curl --header "Authorization: Bearer token"  http://localhost:8080/user
	app.Get("/user", func(ctx fiber.Ctx) error {
		// jwt验证通过后，会把jwt.MapClaims放入Locals
		userInfo := ctx.Locals("user")
		return response.New(ctx).JSON(userInfo)
	})
}

// jwt密钥
const secret = "my-jwt-secret"

type User struct {
	Name     string `validate:"required,min=5,max=20"`
	Password string `validate:"required,min=5,max=20"`
}
