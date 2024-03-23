package server

import (
	"github.com/gofiber/fiber/v3"
)

// New server
func New(router func(app *fiber.App), config ...Config) (err error) {
	cfg := configDefault(config...) // Init config
	app := fiber.New(*cfg.Fiber)
	for _, middleware := range cfg.Middleware {
		app.Use(middleware)
	}

	router(app)
	return app.Listen(cfg.Addr, cfg.ListenConfig)
}
