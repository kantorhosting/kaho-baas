package handlers

import (
	"Kaho_BaaS/internal/apps/account/services"
	"Kaho_BaaS/internal/pkg/sessionmanager"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AccountHandler interface {
	AccountHomeHandler(c *fiber.Ctx) error
	LoginHandler(c *fiber.Ctx) error
	RegisterHandler(c *fiber.Ctx) error
}

type accountHandler struct {
	service services.AccountService
	DB      *gorm.DB
	session *sessionmanager.SessionManager
}

func NewAccountHandler(service services.AccountService, db *gorm.DB, sessionManager *sessionmanager.SessionManager) AccountHandler {
	return &accountHandler{
		service: service,
		DB:      db,
		session: sessionManager,
	}
}
