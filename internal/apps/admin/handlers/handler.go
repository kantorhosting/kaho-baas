package handlers

import (
	"Kaho_BaaS/internal/apps/admin/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AdminHandler interface {
	AdminHomeHandler(c *fiber.Ctx) error
}

type adminHandler struct {
	service services.AdminService
	DB      *gorm.DB
}

func NewAdminHandler(service services.AdminService, db *gorm.DB) AdminHandler {
	return &adminHandler{
		service: service,
		DB:      db,
	}
}
