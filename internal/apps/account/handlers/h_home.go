package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (s *accountHandler) AccountHomeHandler(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	users, err := s.service.FindUsers(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.JSON(fiber.Map{"users": users})
}
