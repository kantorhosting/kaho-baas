package handlers

import "github.com/gofiber/fiber/v2"

func (s *adminHandler) AdminHomeHandler(c *fiber.Ctx) error {
	return c.SendString("Admin Home")
}
