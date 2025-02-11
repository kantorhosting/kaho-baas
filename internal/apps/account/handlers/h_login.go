package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *accountHandler) LoginHandler(c *fiber.Ctx) error {
	projectID := c.Get("X-Kaho-Project") // Ambil project ID dari header
	if projectID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Project ID is required"})
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	// Simulasi autentikasi user project
	if email == "user@project.com" && password == "user123" {
		sess, err := h.session.GetSessionInstance(projectID).Get(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Session error"})
		}

		// Simpan session
		sess.Set("user_id", 1001)
		sess.Set("project_id", projectID)
		sess.Save()

		return c.JSON(fiber.Map{"message": "Login berhasil", "project": projectID})
	}

	return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
}
