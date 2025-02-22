package handlers

import (
	"Kaho_BaaS/internal/apps/account/models"
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// RegisterHandler godoc
//
//	@Summary		Register user for a project
//	@Description	Authenticate user credentials and start a user session.
//	@Tags			account
//	@Accept			application/x-www-form-urlencoded, json
//	@Produce		json
//	@Param			X-Kaho-Project	header		string					true	"Project ID"
//	@Param			name			formData	string					true	"User Name"
//	@Param			email			formData	string					true	"User Email"
//	@Param			password		formData	string					true	"User Password"
//	@Param			confirmPassword	formData	string					true	"Confirm Password"
//	@Success		201				{object}	models.Session			"Register success response"
//	@Failure		400				{object}	map[string]string		"X-Kaho-Project is required"
//	@Failure		400				{object}	map[string]string		"Password unmatch with confirm password"
//	@Failure		400				{object}	map[string]string		"Register failed"
//	@Failure		500				{object}	map[string]interface{}	"Server error"
//	@Router			api/v1/register [post]
func (h *accountHandler) RegisterHandler(c *fiber.Ctx) error {
	//NOTE: use 2 secs because got timeout when using 1 sec
	ctx, cancel := context.WithTimeout(c.UserContext(), 2*time.Second)
	defer cancel()

	projectID := c.Get("X-Kaho-Project")
	if projectID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "X-Kaho-Project is required"})
	}

	data := new(models.Register)
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

	user, token, err := h.service.Register(ctx, data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	sess, err := h.session.GetSessionInstance(projectID).Get(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Session error"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Token",
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	sess.Set("user_id", user.ID)
	sess.Set("project_id", projectID)
	sess.Save()

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Register Success", "project": projectID, "user": user})
}
