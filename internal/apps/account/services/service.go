package services

import (
	"Kaho_BaaS/internal/apps/account/models"
	"Kaho_BaaS/internal/apps/account/repositories"
	"Kaho_BaaS/internal/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type AccountService interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	Register(ctx context.Context, data *models.Register) (*models.User, error)
	Login(ctx context.Context, data *models.Login) (*models.User, error)
}

type accountService struct {
	repository repositories.AccountRepository
}

func NewAccountService(repository repositories.AccountRepository) AccountService {
	return &accountService{
		repository: repository,
	}
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
func (as *accountService) Register(ctx context.Context, data *models.Register) (*models.User, error) {
	_, err := as.repository.FindUserByEmail(ctx, data.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("User already exists",
			"email", data.Email,
			"err", err,
		)

		return nil, fmt.Errorf("User already exists")
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		slog.Error("Failed hashing password",
			"err", err,
		)

		return nil, fmt.Errorf("Unexpected error happened. Please try again!")
	}

	data.Password = hashedPassword
	user, err := as.repository.Create(ctx, data)
	if err != nil {
		slog.Error("Failed creating user",
			"err", err,
		)

		return nil, fmt.Errorf("Unexpected error happened. Please try again!")
	}

	return user, nil
}

// Login implements AccountService.
func (as *accountService) Login(ctx context.Context, data *models.Login) (*models.User, error) {
	user, err := as.repository.FindUserByEmail(ctx, data.Email)
	if err != nil {
		slog.Error("Failed retrieve user",
			"email", data.Email,
			"err", err,
		)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User not found")
		}

		return nil, fmt.Errorf("Unexpected error happened. Please try again!")
	}

	isMatch := utils.CheckPasswordHash(data.Password, user.Password)
	if !isMatch {
		return nil, fmt.Errorf("Invalid credentials")
	}

	return user, nil
}
