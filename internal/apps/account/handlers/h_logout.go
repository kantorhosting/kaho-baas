package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// LogoutHandler godoc
//
//	@Summary		Logout user from a project
//	@Description	End the user session by destroying the session data.
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			X-Kaho-Project	header		string					true	"Project ID"
//	@Success		200				{object}	map[string]string		"Logout success response"
//	@Failure		400				{object}	map[string]string		"X-Kaho-Project is required"
//	@Failure		500				{object}	map[string]string		"Failed to destroy session"
//	@Router			/account/sessions/logout [delete]
func (h *accountHandler) LogoutHandler(c *fiber.Ctx) error {
	projectID := c.Get("X-Kaho-Project")
	if projectID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "X-Kaho-Project is required"})
	}
	sess, err := h.session.GetSessionInstance(projectID).Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve session"})
	}

	if err := sess.Destroy(); err != nil {
		slog.Error("Failed to destroy session", "err", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to logout"})
	}

	return c.JSON(fiber.Map{"message": "Logout Success"})

}
