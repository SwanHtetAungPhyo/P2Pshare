package main

import (
	"log"

	"github.com/SwanHtetAungPhyo/api_gate/internal/config"
	"github.com/SwanHtetAungPhyo/api_gate/internal/logging"
	"github.com/goccy/go-json"
	"go.uber.org/zap"

	"github.com/SwanHtetAungPhyo/api_gate/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}
	logging.InitLogger()
	defer logging.CloseLogger()

	port := cfg.ServiceLeader.Env

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(middleware.LoggingMetrics())
	middleware.ServiceRegistry(cfg, app)
	logging.Logger.Info("gateway is listening on the server: port: 8081",

		zap.String("service", cfg.ServiceLeader.Name),
	)
	if app := app.Listen(port[0].Value); app != nil {
		log.Fatal(app)
	}
}
