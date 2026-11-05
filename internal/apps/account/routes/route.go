package routes

import (
	"Kaho_BaaS/internal/apps/account/handlers"
	"Kaho_BaaS/internal/apps/account/repositories"
	"Kaho_BaaS/internal/apps/account/services"
	"Kaho_BaaS/internal/pkg/middlewares"
	"Kaho_BaaS/internal/pkg/sessionmanager"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, sessionManager *sessionmanager.SessionManager) {
	accountRepository := repositories.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(accountService, sessionManager)

	router.Post("/login", accountHandler.LoginHandler)
	router.Post("/register", accountHandler.RegisterHandler)

	accountGroup := router.Group("/account")
	accountGroup.Use(middlewares.AuthOnly)

	accountGroup.Get("/", accountHandler.AccountHomeHandler)
}
