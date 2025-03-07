package middleware

import (
	"time"

	"github.com/SwanHtetAungPhyo/api_gate/internal/logging"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func LoggingMetrics() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()
		url := c.Request().URI().String()

		logging.Logger.Info("Request completed",
			zap.String("URL", url),
			zap.Duration("Duration", time.Since(start)),
			zap.String("Method", c.Method()),
			zap.Int("Status", c.Response().StatusCode()),
			zap.String("User-Agent", c.Get("User-Agent")),
			zap.String("IP", c.IP()),
			zap.String("Referer", c.Get("Referer")),
			zap.String("Origin", c.Get("Origin")),
			zap.String("Host", c.Get("Host")),
			zap.String("Response", string(c.Response().Body())),
		)

		return c.Next()
	}
}
