package handlers

import (
	"Kaho_BaaS/internal/apps/account/models"
	"context"
	"log/slog"
	"net/http"
	"time"

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
//	@Success		200				{object}	string			        "Login success response"
//	@Failure		400				{object}	map[string]string		"X-Kaho-Project is required"
//	@Failure		401				{object}	map[string]string		"Invalid credentials"
//	@Failure		500				{object}	map[string]interface{}	"Server error"
//	@Router			/account/sessions/login [post]
func (h *accountHandler) LoginHandler(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.UserContext(), 1*time.Second)
	defer cancel()

	projectID := c.Get("X-Kaho-Project")
	if projectID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "X-Kaho-Project is required"})
	}

	data := new(models.Login)
	if err := c.BodyParser(data); err != nil {
		slog.Error("Failed parsing request body",
			"err", err,
		)

		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Request body invalid"})
	}

	if errs := h.validator.Validate(data); errs != nil && len(errs) > 0 {
		slog.Error("Request body contain invalid data")

		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": errs})
	}

	user, err := h.service.Login(ctx, data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	sess, err := h.session.GetSessionInstance(projectID).Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Session error"})
	}

	// Simpan session
	sess.Set("user_id", user.ID)
	sess.Set("project_id", projectID)
	sess.Save()

	return c.JSON(fiber.Map{"message": "Login Success", "project": projectID})
}
