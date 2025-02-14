package database

import (
	"Kaho_BaaS/internal/apps/account/models"
	"log/slog"

	"gorm.io/gorm"
)

var modelList = map[string]interface{}{
	"users": &models.User{},
}

func autoMigrateModels(DB *gorm.DB) {
	for name, model := range modelList {
		slog.Info("Migrating start for",
			"entity", name,
		)

		err := DB.AutoMigrate(model)
		if err != nil {
			slog.Error("Migrating error for",
				"entity", name,
				"err", err,
			)
			continue
		}

		slog.Info("Migrating done for",
			"entity", name,
		)
	}
}
