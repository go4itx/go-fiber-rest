package server

import (
	"log"

	"github.com/go4itx/go-fiber-rest/response"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

type Config struct {
	AppName      string             // 默认："go-fiber-rest"
	Addr         string             // 默认："0.0.0.0:8080"
	ListenConfig fiber.ListenConfig //
	Fiber        *fiber.Config
	Middleware   []fiber.Handler
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Addr:    "0.0.0.0:8080",
	AppName: "go-fiber-rest",
	Fiber: &fiber.Config{
		BodyLimit: 4 * 1024 * 1024, // 4MB
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			return response.New(ctx).JSON(err)
		},
	},
	Middleware: []fiber.Handler{
		recover.New(recover.Config{
			EnableStackTrace: true,
			StackTraceHandler: func(ctx fiber.Ctx, e any) {
				log.Println("========StackTrace========")
				log.Println(e)
			},
		}),
		logger.New(),
		// cors.New(),
		// cors.New(cors.Config{
		// 	AllowOrigins:     "http://0.0.0.0:8080",
		// 	AllowCredentials: true,
		// }),
	},
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Addr == "" {
		cfg.Addr = ConfigDefault.Addr
	}

	if cfg.Fiber == nil {
		cfg.Fiber = ConfigDefault.Fiber
	}

	if len(cfg.Middleware) < 1 {
		cfg.Middleware = ConfigDefault.Middleware
	}

	if cfg.AppName != "" {
		cfg.Fiber.AppName = cfg.AppName
	}

	return cfg
}
