package migration

import (
	"gogo/config"
	"gogo/model"
)

func CreateMigration() {
	config.DB.AutoMigrate(&model.Migration{})
}