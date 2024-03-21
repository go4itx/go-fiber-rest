package server

import (
	"github.com/gofiber/fiber/v3"
)

// Init server
func New(router func(app *fiber.App), config ...Config) (err error) {
	// Init config
	cfg := configDefault(config...)
	app := fiber.New(*cfg.Fiber)
	for _, middleware := range cfg.Middleware {
		app.Use(middleware)
	}

	router(app)
	return app.Listen(cfg.Addr, cfg.ListenConfig)
}
