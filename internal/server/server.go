package server

import (
	"github.com/gofiber/fiber/v2"

	"Kaho_BaaS/internal/database"
	"Kaho_BaaS/internal/pkg/sessionmanager"
)

type FiberServer struct {
	*fiber.App

	db             database.Service
	gormDB         database.ServiceGorm
	sessionmanager *sessionmanager.SessionManager
}

func New() *FiberServer {
	gormDB, err := database.ConnectDatabase()
	if err != nil {
		return nil
	}

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader:  "Kaho_BaaS",
			AppName:       "Kaho_BaaS",
			StrictRouting: true,
			CaseSensitive: true,
		}),

		db:             database.New(),
		gormDB:         gormDB,
		sessionmanager: sessionmanager.NewSessionManager(),
	}

	return server
}
