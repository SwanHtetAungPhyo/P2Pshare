package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/SwanHtetAungPhyo/api_gate/internal/config"
	"github.com/SwanHtetAungPhyo/api_gate/internal/logging"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
)

var cbInstance *gobreaker.CircuitBreaker

func init() {
	cbInstance = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "Cloumbs",
		MaxRequests: 5,
		Interval:    10 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
	})
}
func ServiceRegistry(cfg *config.Config, app *fiber.App) {
	for _, svc := range cfg.ServiceLeader.Services {
		serviceGroup := app.Group(svc.Prefix)

		serviceGroup.Add(svc.AllowedMethods, "/*", func(target string) fiber.Handler {
			return func(c fiber.Ctx) error {
				response, err := cbInstance.Execute(func() (interface{}, error) {
					uri := c.OriginalURL()
					backendURL := GetBackendURL(uri, target)

					logging.Logger.Info("Forwarding request",
						zap.String("Backend URL", backendURL),
					)

					err := proxy.Do(c, backendURL)
					if err != nil {
						logging.Logger.Error("Service call failed", zap.Error(err))
						return nil, errors.New("service unavailable")
					}
					return nil, nil
				})

				if err != nil {
					return c.Status(fiber.StatusServiceUnavailable).JSON(
						fiber.Map{"error": "Service temporarily unavailable"},
					)
				}
				return response.(error)
			}
		}(svc.URLs[0]))

		logging.Logger.Debug("Service Routed",
			zap.String("Name", svc.Name),
			zap.Strings("Methods", svc.AllowedMethods),
			zap.String("Prefix", svc.Prefix),
		)
	}
	WareSetup(cfg, app)
}

func GetBackendURL(uri string, target string) string {
	parts := strings.Split(uri, "/")
	lastParts := parts[len(parts)-1]
	return target + "/" + lastParts
}
