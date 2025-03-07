package route

import (
	"github.com/SwanHtetAungPhyo/auth-service/internal/handler"
	"github.com/SwanHtetAungPhyo/auth-service/internal/repo"
	"github.com/SwanHtetAungPhyo/auth-service/internal/sevice"
	"github.com/gofiber/fiber/v2"
)

func RouteSetup(app *fiber.App) {
	repo := repo.NewUserRepo()
	authService := sevice.NewAuthService(*repo)
	authHandler := handler.NewAuthHandler(authService)
	app.Post("/register", authHandler.Register)
}
