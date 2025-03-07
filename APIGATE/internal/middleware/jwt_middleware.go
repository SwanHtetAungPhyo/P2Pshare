package middleware

import (
	"fmt"
	"strings"

	"github.com/SwanHtetAungPhyo/api_gate/internal/config"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JwtMiddleware(cfg *config.Config) fiber.Handler {
	return func(c fiber.Ctx) error {
		requestedPath := c.OriginalURL()
		for _, svc := range cfg.ServiceLeader.Services {
			if strings.HasPrefix(requestedPath, svc.Prefix) && shouldSkipJWT(svc) {
				return c.Next()
			}
		}
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"error": "Unauthorized",
				})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"error": "Invalid token format",
				})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.ServiceLeader.Env[2].Value), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"error": "Invalid token",
				})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"error": "Invalid token",
				})
		}
		return c.Next()
	}
}

func shouldSkipJWT(svc config.Service) bool {
	if skip, ok := svc.Filter["skip_jwt"].(bool); ok && skip {
		return true
	}
	return false
}
