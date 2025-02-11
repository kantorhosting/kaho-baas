package handlers

import (
	"Kaho_BaaS/internal/apps/admin/models"

	"github.com/gofiber/fiber/v2"
)

func (s *adminHandler) AdminHomeHandler(c *fiber.Ctx) error {
	var users []models.User
	s.DB.AutoMigrate(&models.User{})
	s.DB.AutoMigrate(&models.Target{})
	s.DB.AutoMigrate(&models.HashOptions{})
	result := s.DB.Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	return c.JSON(fiber.Map{
		"users": users,
	})
}
