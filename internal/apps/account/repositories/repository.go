package repositories

import (
	"Kaho_BaaS/internal/apps/account/models"
	"context"

	"gorm.io/gorm"
)

type AccountRepository interface {
	FindUsers(ctx context.Context) ([]models.User, error)
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

// FindUsers implements AccountRepository.
func (as *accountRepository) FindUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	result := as.db.WithContext(ctx).Debug().Find(&users)

	return users, result.Error
}

// FindUserByEmail implements AccountService.
func (as *accountRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	result := as.db.WithContext(ctx).Debug().Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
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
