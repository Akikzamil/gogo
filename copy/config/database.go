package config

import (
	"fmt"
	"gogo/model"
	"gogo/utils"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB;
var err error

func InitializeDatabaseConnection() bool {
	dbConnection := utils.Getenv("DB_CONNECTION", "mysql")
	dbHost := utils.Getenv("DB_HOST", "127.0.0.1")
	dbPort := utils.Getenv("DB_PORT", "3306")
	dbDatabase := utils.Getenv("DB_DATABASE", "3306")
	dbUser := utils.Getenv("DB_USERNAME", "root")
	dbPass := utils.Getenv("DB_PASSWORD", "")


	switch dbConnection {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDatabase)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		  });
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbDatabase, dbPort);
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		  });
	}

	return err == nil;
}


func RunAllMigrations() {
	DB.AutoMigrate(&model.Migration{})
}
