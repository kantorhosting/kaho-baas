package routes

import (
	"Kaho_BaaS/internal/apps/account/handlers"
	"Kaho_BaaS/internal/apps/account/services"
	"Kaho_BaaS/internal/pkg/sessionmanager"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, sessionManager *sessionmanager.SessionManager) {
	accountGroup := router.Group("/account")
	accountService := services.NewAccountService()
	accountHandler := handlers.NewAccountHandler(accountService, db, sessionManager)

	accountGroup.Get("/", accountHandler.AccountHomeHandler)
	accountGroup.Post("/sessions/email", accountHandler.LoginHandler)
}
