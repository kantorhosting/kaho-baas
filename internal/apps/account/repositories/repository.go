package repositories

import (
	"Kaho_BaaS/internal/apps/account/models"
	"context"

	"gorm.io/gorm"
)

type AccountRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, data *models.Register) (*models.User, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

// FindUserByEmail implements AccountService.
func (as *accountRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	result := as.db.WithContext(ctx).Debug().First(&user)

	return &user, result.Error
}

// Create implements AccountService.
func (as *accountRepository) Create(ctx context.Context, data *models.Register) (*models.User, error) {
	user := models.User{
		Email:    data.Email,
		Password: data.Password,
		Name:     data.Name,
	}

	result := as.db.WithContext(ctx).Debug().Create(&user)

	return &user, result.Error
}
