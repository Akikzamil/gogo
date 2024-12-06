package main

import (
	"fmt"
	"gogo/config"
	"gogo/route"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	config.InitializeDatabaseConnection()

	var commandName string
	if len(os.Args) > 1 {
		commandName = os.Args[1]
	}

	switch commandName {
	case "migrate":
		config.RunAllMigrations();
	default:
		run();
	}
}

func run() {
	app := fiber.New()

	route.SetUpRoutes(app)

	port := getPort()
	app.Listen(fmt.Sprintf(":" + port))
}

func getPort() string {
	port := ""
	err := godotenv.Load()

	if err != nil {
		port = "3000"
	} else {
		envPort := os.Getenv("PORT")
		if envPort != "" {
			port = envPort
		} else {
			port = "3000"
		}
	}
	return port
}
