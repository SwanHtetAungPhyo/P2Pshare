package middleware

import (
	"time"

	"github.com/SwanHtetAungPhyo/api_gate/internal/config"
	"github.com/SwanHtetAungPhyo/api_gate/internal/logging"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"go.uber.org/zap"
)

func WareSetup(cfg *config.Config, app *fiber.App) {
	for key, value := range cfg.ServiceLeader.Filter { // FIX: Directly iterate over map
		switch key {
		case "jwt":
			if enable, ok := value.(bool); ok && enable {
				app.Use(JwtMiddleware(cfg))
				logging.Logger.Info("JWT middleware set up",
					zap.String("service", cfg.ServiceLeader.Name),
					zap.String("version", cfg.ServiceLeader.Version),
				)
			}
		case "rate_limit":
			if limit, ok := value.(int); ok && limit > 0 {
				app.Use(limiter.New(limiter.Config{
					Max:        limit,
					Expiration: 1 * time.Minute,
				}))
			}
		}
	}
}
