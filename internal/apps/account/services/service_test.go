package services

import (
	"Kaho_BaaS/internal/apps/account/models"
	"Kaho_BaaS/internal/apps/account/repositories"
	"Kaho_BaaS/internal/pkg/constants"
	"Kaho_BaaS/internal/pkg/utils"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var repository = &repositories.AccountRepositoryMock{Mock: mock.Mock{}}
var service = accountService{repository: repository}

func TestFindByEmail_NotFound(t *testing.T) {
	email := "test123@demo.com"
	repository.Mock.On("FindUserByEmail", email).Return(nil)

	user, err := service.FindUserByEmail(context.TODO(), email)

	assert.NotNil(t, err)
	assert.EqualValues(t, gorm.ErrRecordNotFound, err)
	assert.Nil(t, user)
}

func TestFindByEmail_Found(t *testing.T) {
	email := "john123@demo.com"
	repository.Mock.On("FindUserByEmail", email).Return(models.User{
		Email: email,
	})

	user, err := service.FindUserByEmail(context.TODO(), email)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
}

func TestLogin_NotFound(t *testing.T) {
	data := models.Login{
		Email:    "abc123@demo.com",
		Password: "john123!@#",
	}

	repository.Mock.On("FindUserByEmail", data.Email).Return(nil)

	user, err := service.Login(context.TODO(), &data)

	assert.NotNil(t, err)
	assert.EqualValues(t, constants.ErrUserNotFound, err)
	assert.Nil(t, user)
}

func TestLogin_InvalidCred(t *testing.T) {
	data := models.Login{
		Email:    "john123@demo.com",
		Password: "john123!@#",
	}

	repository.Mock.On("FindUserByEmail", data.Email).Return(models.User{
		Email: data.Email,
	})

	user, err := service.Login(context.TODO(), &data)

	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("Invalid credentials"), err)
	assert.Nil(t, user)

}

func TestLogin_Success(t *testing.T) {
	data := models.Login{
		Email:    "johndoes123@demo.com",
		Password: "john123!@#",
	}

	hashedPassword, _ := utils.HashPassword(data.Password)
	repository.Mock.On("FindUserByEmail", data.Email).Return(models.User{
		Email:    data.Email,
		Password: hashedPassword,
	})

	user, err := service.Login(context.TODO(), &data)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, data.Email, user.Email)
}

func TestRegister_UserAlreadyExist(t *testing.T) {
	data := models.Register{
		Email:    "john123@demo.com",
		Password: "john123!@#",
	}

	repository.Mock.On("FindUserByEmail", data.Email).Return(models.User{
		Email: data.Email,
	})
	repository.Mock.On("Create", data).Return(nil)

	user, err := service.Register(context.TODO(), &data)

	assert.NotNil(t, err)
	assert.EqualValues(t, constants.ErrUserAlreadyExist, err)
	assert.Nil(t, user)
}

func TestRegister_PasswordTooLong(t *testing.T) {
	data := models.Register{
		Email:    "johndoe123@demo.com",
		Password: "john123!@#awdjwaiodjawodjawodjawodiawjdioawjdoiawjdioawjdioawjdiojaslkdhasjifghasjifghasjkfhsajkd",
	}

	repository.Mock.On("FindUserByEmail", data.Email).Return(nil)
	repository.Mock.On("Create", data).Return(nil)

	user, err := service.Register(context.TODO(), &data)

	assert.NotNil(t, err)
	assert.EqualValues(t, constants.ErrInternalServer, err)
	assert.Nil(t, user)
}

func TestRegister_Success(t *testing.T) {
	data := models.Register{
		Email:    "johndoe123@demo.com",
		Password: "johndoes123!@#",
	}

	repository.Mock.On("FindUserByEmail", data.Email).Return(nil)
	repository.Mock.On("Create", data).Return(models.User{
		Email: data.Email,
	})

	user, err := service.Register(context.TODO(), &data)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, data.Email, user.Email)
}
