package server

import (
	"github.com/gofiber/fiber/v2"

	"Kaho_BaaS/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "Kaho_BaaS",
			AppName:      "Kaho_BaaS",
		}),

		db: database.New(),
	}

	return server
}
