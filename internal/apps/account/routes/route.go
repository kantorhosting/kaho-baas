package routes

import (
	"Kaho_BaaS/internal/apps/account/handlers"
	"Kaho_BaaS/internal/apps/account/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	accountGroup := router.Group("/account")
	accountService := services.NewAccountService()
	accountHandler := handlers.NewAccountHandler(accountService, db)

	accountGroup.Get("/", accountHandler.AccountHomeHandler)
}
