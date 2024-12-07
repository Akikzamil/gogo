package migration

import (
	"gogo/config"
	"gogo/model"
)

func (m MigrtionStruct) Up() error {
	err := config.DB.Migrator().CreateTable(&model.GOGO{})
	return err;
}

func (m MigrtionStruct) Down() error {
	err := config.DB.Migrator().DropTable(&model.GOGO{})
	return err;
}