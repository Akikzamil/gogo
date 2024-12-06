package migration

import (
	"gogo/config"
	"gogo/model"
)

func Up() {
	config.DB.Migrator().CreateTable(&model.GOGO{})
}

func Down() {
	config.DB.Migrator().DropTable(&model.GOGO{})
}