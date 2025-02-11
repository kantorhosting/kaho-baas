package handlers

import (
	"Kaho_BaaS/internal/apps/account/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AccountHandler interface {
	AccountHomeHandler(c *fiber.Ctx) error
}

type accountHandler struct {
	service services.AccountService
	DB      *gorm.DB
}

func NewAccountHandler(service services.AccountService, db *gorm.DB) AccountHandler {
	return &accountHandler{
		service: service,
		DB:      db,
	}
}
