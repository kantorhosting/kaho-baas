package middlewares

import (
	"Kaho_BaaS/internal/apps/account/models"
	"Kaho_BaaS/internal/pkg/utils"
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthOnly(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token")
	if tokenStr == "" {
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(strings.Split(authHeader, " ")) == 0 {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authentication failed. Please sign in first"})
		}

		tokenStr = strings.Split(authHeader, " ")[1]
	}

	token, err := utils.VerifyingJWT(os.Getenv("JWT_SECRET"), tokenStr)
	if err != nil {
		slog.Error("Verifying JWT failed",
			"err", err,
		)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authentication failed. Please sign in first"})
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		slog.Error("Extract JWT claims failed",
			"err", err,
		)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authentication failed. Please sign in first"})
	}

	c.SetUserContext(context.WithValue(c.UserContext(), "user", claims))
	return c.Next()
}
