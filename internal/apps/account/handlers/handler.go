package handlers

import (
	"Kaho_BaaS/internal/apps/account/services"
	"Kaho_BaaS/internal/pkg/sessionmanager"
	"Kaho_BaaS/internal/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler interface {
	AccountHomeHandler(c *fiber.Ctx) error
	LoginHandler(c *fiber.Ctx) error
	RegisterHandler(c *fiber.Ctx) error
}

type accountHandler struct {
	service   services.AccountService
	session   *sessionmanager.SessionManager
	validator *utils.Validator
}

func NewAccountHandler(service services.AccountService, sessionManager *sessionmanager.SessionManager) AccountHandler {
	return &accountHandler{
		service:   service,
		session:   sessionManager,
		validator: utils.NewValidator(),
	}
}
