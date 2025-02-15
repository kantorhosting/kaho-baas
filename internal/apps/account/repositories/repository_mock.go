package repositories

import (
	"Kaho_BaaS/internal/apps/account/models"
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type AccountRepositoryMock struct {
	Mock mock.Mock
}

func (ar *AccountRepositoryMock) FindUsers(ctx context.Context) ([]models.User, error) {
	return []models.User{}, nil
}

func (ar *AccountRepositoryMock) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	args := ar.Mock.Called(email)

	if args.Get(0) == nil {
		return nil, gorm.ErrRecordNotFound
	}

	user := args.Get(0).(models.User)

	return &user, nil
}

func (ar *AccountRepositoryMock) Create(ctx context.Context, data *models.Register) (*models.User, error) {
	user := models.User{
		Email:    data.Email,
		Password: data.Password,
		Name:     data.Name,
	}

	return &user, nil
}
