package routes

import (
	"Kaho_BaaS/internal/apps/admin/handlers"
	"Kaho_BaaS/internal/apps/admin/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	adminGroup := router.Group("/admin")
	adminService := services.NewAdminService()
	adminHandler := handlers.NewAdminHandler(adminService)

	adminGroup.Get("/", adminHandler.AdminHomeHandler)
}
