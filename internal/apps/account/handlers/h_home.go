package handlers

import (
	"Kaho_BaaS/internal/apps/account/models"

	"github.com/gofiber/fiber/v2"
)

func (s *accountHandler) AccountHomeHandler(c *fiber.Ctx) error {
	var users []models.User
	result := s.DB.Select("\"$id\", name, email, \"$createdAt\", \"$updatedAt\"").Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	return c.JSON(fiber.Map{
		"users": users,
	})
}
