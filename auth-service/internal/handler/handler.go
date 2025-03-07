package handler

import (
	"github.com/SwanHtetAungPhyo/auth-service/internal/model"
	"github.com/SwanHtetAungPhyo/auth-service/internal/sevice"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *sevice.AuthService
}

func NewAuthHandler(authService *sevice.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var modelUser *model.User
	if err := c.BodyParser(&modelUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := h.authService.Register(modelUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&Response{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&Response{
		Status:  "success",
		Message: "User registered successfully",
		Data:    modelUser,
	})
}
