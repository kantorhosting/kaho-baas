package services

import (
	"Kaho_BaaS/internal/apps/account/models"
	"Kaho_BaaS/internal/apps/account/repositories"
	"Kaho_BaaS/internal/pkg/constants"
	"Kaho_BaaS/internal/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"gorm.io/gorm"
)

type AccountService interface {
	FindUsers(ctx context.Context) ([]models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	Register(ctx context.Context, data *models.Register) (*models.User, string, error)
	Login(ctx context.Context, data *models.Login) (*models.User, string, error)
}

type accountService struct {
	repository repositories.AccountRepository
}

func NewAccountService(repository repositories.AccountRepository) AccountService {
	return &accountService{
		repository: repository,
	}
}

// FindUsers implements AccountService.
func (as *accountService) FindUsers(ctx context.Context) ([]models.User, error) {
	users, err := as.repository.FindUsers(ctx)
	if err != nil {
		slog.Error("Retrieve all users",
			"err", err,
		)

		return []models.User{}, err
	}

	return users, nil
}

// FindUserByEmail implements AccountService.
func (as *accountService) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := as.repository.FindUserByEmail(ctx, email)
	if err != nil {
		slog.Error("Retrieve user",
			"email", email,
			"err", err,
		)
		return nil, err
	}

	return user, nil
}

// Create implements AccountService.
func (as *accountService) Register(ctx context.Context, data *models.Register) (*models.User, string, error) {
	if data.Password != data.ConfirmPassword {
		slog.Error("Password unmatch with confirm password")

		return nil, "", fmt.Errorf("Password unmatch with confirm password")
	}

	user, err := as.repository.FindUserByEmail(ctx, data.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed retrieve user",
			"email", data.Email,
			"err", err,
		)

		return nil, "", constants.ErrInternalServer
	}

	if user != nil {
		slog.Error("User already exist",
			"email", user.Email)

		return nil, "", constants.ErrUserAlreadyExist
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		slog.Error("Failed hashing password",
			"err", err,
		)

		return nil, "", constants.ErrInternalServer
	}

	data.Password = hashedPassword
	user, err = as.repository.Create(ctx, data)
	if err != nil {
		slog.Error("Failed creating user",
			"err", err,
		)

		return nil, "", constants.ErrInternalServer
	}

	token, err := utils.GenerateJWT(os.Getenv("JWT_SECRET"), user.ID, user.Email, user.Name)
	if err != nil {
		slog.Error("Failed generating token",
			"err", err,
		)

		return nil, "", fmt.Errorf("Your data has been saved but failed generating auth token. Please try signin with your previous data")
	}

	return user, token, nil
}

// Login implements AccountService.
func (as *accountService) Login(ctx context.Context, data *models.Login) (*models.User, string, error) {
	user, err := as.repository.FindUserByEmail(ctx, data.Email)
	if err != nil {
		slog.Error("Failed retrieve user",
			"email", data.Email,
			"err", err,
		)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", constants.ErrUserNotFound
		}

		return nil, "", constants.ErrInternalServer
	}

	isMatch := utils.CheckPasswordHash(data.Password, user.Password)
	if !isMatch {
		return nil, "", fmt.Errorf("Invalid credentials")
	}

	token, err := utils.GenerateJWT(os.Getenv("JWT_SECRET"), user.ID, user.Email, user.Name)
	if err != nil {
		slog.Error("Failed generating token",
			"err", err,
		)

		return nil, "", fmt.Errorf("Your data has been saved but failed generating auth token. Please try signin with your previous data")
	}

	return user, token, nil
}
