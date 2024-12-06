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
	app := fiber.New();
	config.InitializeDatabaseConnection();

	route.SetUpRoutes(app);

	port := getPort();
	app.Listen(fmt.Sprintf(":"+port))
}


func getPort() string{
	port:= "";
	err := godotenv.Load();

	if err != nil {
		port = "3000"
	}else{
	 envPort := os.Getenv("PORT");
	 if(envPort!="") {
		port = envPort
	 }else{
		port = "3000"
	 }
	}
	return port
}