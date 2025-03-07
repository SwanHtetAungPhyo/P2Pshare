package sevice

import (
	"context"

	"github.com/SwanHtetAungPhyo/auth-service/internal/logging"
	"github.com/SwanHtetAungPhyo/auth-service/internal/model"
	"github.com/SwanHtetAungPhyo/auth-service/internal/repo"
	"github.com/SwanHtetAungPhyo/protos/user"
	"go.uber.org/zap"
)

type AuthService struct {
	repo repo.UserRepo
}

func NewAuthService(repo repo.UserRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) Register(user *model.User) error {
	logging.Logger.Info("Registering the new User")
	if err := a.repo.Create(user); err != nil {
		logging.Logger.Error("Error in registering ", zap.Error(err))
		return err
	}

	logging.Logger.Info("User registered successfully")
	return nil
}

type UserServiceServerImpl struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServiceServerImpl) CheckUserExists(ctx context.Context, req *user.CheckUserExistsRequest) (*user.CheckUserExistsResponse, error) {
	// Implement your logic to check if the user exists
	exists := false // Replace with actual check
	return &user.CheckUserExistsResponse{Exists: exists}, nil
}
