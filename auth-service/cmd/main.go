package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/SwanHtetAungPhyo/auth-service/internal/logging"
	"github.com/SwanHtetAungPhyo/auth-service/internal/route"
	"github.com/SwanHtetAungPhyo/auth-service/internal/sevice"
	"github.com/SwanHtetAungPhyo/protos/user"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	grpcServer := grpc.NewServer()
	logging.InitLogger()

	lis, err := net.Listen("tcp", ":50051")
	logging.Logger.Info("Grpc server is starting")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	route.RouteSetup(app)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()
	userServiceServer := &sevice.UserServiceServerImpl{}
	user.RegisterUserServiceServer(grpcServer, userServiceServer)
	go func() {
		if err := app.Listen(":8001"); err != nil {
			panic(err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	log.Printf("Received signal: %v. Shutting down servers...", sig)

	// Gracefully stop the servers
	gracefulShutdown(grpcServer, app)
}

func gracefulShutdown(grpcServer *grpc.Server, app *fiber.App) {
	// Create a context with a timeout for graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stop the Fiber server gracefully
	if err := app.Shutdown(); err != nil {
		log.Printf("Fiber server shutdown error: %v", err)
	}

	// Gracefully stop the gRPC server
	grpcServer.GracefulStop()
	log.Println("Servers stopped gracefully")
}
