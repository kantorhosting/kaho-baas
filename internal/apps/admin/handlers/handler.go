package handlers

import (
	"Kaho_BaaS/internal/apps/admin/services"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler interface {
	AdminHomeHandler(c *fiber.Ctx) error
}

type adminHandler struct {
	service services.AdminService
}

func NewAdminHandler(service services.AdminService) AdminHandler {
	return &adminHandler{
		service: service,
	}
}
