package routes

import (
	"Kaho_BaaS/internal/apps/admin/handlers"
	"Kaho_BaaS/internal/apps/admin/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	adminGroup := router.Group("/admin")
	adminService := services.NewAdminService()
	adminHandler := handlers.NewAdminHandler(adminService, db)

	adminGroup.Get("/", adminHandler.AdminHomeHandler)
}
