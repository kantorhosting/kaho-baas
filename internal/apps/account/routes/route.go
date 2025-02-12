package routes

import (
	"Kaho_BaaS/internal/apps/account/handlers"
	"Kaho_BaaS/internal/apps/account/repositories"
	"Kaho_BaaS/internal/apps/account/services"
	"Kaho_BaaS/internal/pkg/sessionmanager"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, sessionManager *sessionmanager.SessionManager) {
	accountGroup := router.Group("/account")
	accountRepository := repositories.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(accountService, db, sessionManager)

	accountGroup.Get("/", accountHandler.AccountHomeHandler)
	accountGroup.Post("/sessions/login", accountHandler.LoginHandler)
	accountGroup.Post("/sessions/register", accountHandler.RegisterHandler)
}
