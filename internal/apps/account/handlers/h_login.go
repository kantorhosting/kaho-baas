package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// LoginHandler godoc
//
//	@Summary		Login user for a project
//	@Description	Authenticate user credentials and start a user session.
//	@Tags			account
//	@Accept			application/x-www-form-urlencoded
//	@Produce		json
//	@Param			X-Kaho-Project	header		string					true	"Project ID"
//	@Param			email			formData	string					true	"User Email"
//	@Param			password		formData	string					true	"User Password"
//	@Success		200				{object}	models.Session			"Login success response"
//	@Failure		400				{object}	map[string]string		"X-Kaho-Project is required"
//	@Failure		401				{object}	map[string]string		"Invalid credentials"
//	@Failure		500				{object}	map[string]interface{}	"Server error"
//	@Router			/account/sessions/email [post]
func (h *accountHandler) LoginHandler(c *fiber.Ctx) error {
	projectID := c.Get("X-Kaho-Project") // Ambil project ID dari header
	if projectID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "X-Kaho-Project is required"})
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
