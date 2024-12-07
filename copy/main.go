package main

import (
	"fmt"
	"gogo/config"
	"gogo/migration"
	"gogo/model"
	"gogo/route"
	"log"
	"os"
	"reflect"
	"strings"

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
		RunAllMigrations();
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

func RunAllMigrations() {
	migration.CreateMigration();
	var migrationModel1  model.Migration ;
	config.DB.Order("batch DESC").Select("batch").First(&migrationModel1);
	currentBatch := migrationModel1.Batch+1;

	dirPath := "./migration"

	// Read all files in the directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	// Iterate over the files and print their names
	for _, file := range files {
		if(file.Name() == "create_migration.go" || file.Name() == "mirgration.struct.go"){
			continue;
		}
		var mig model.Migration
		config.DB.First(&mig).Where("migration = ?", file.Name);

		if mig.Migration != "" {
			continue
		}
		
		st:= migration.MigrtionStruct{}
		splittedFileName:= strings.Split(file.Name(),".")[0];
		fileName := "Up"+ splittedFileName;
		fmt.Println(file.Name())
		fmt.Println(fileName)
		method := reflect.ValueOf(st).MethodByName(fileName);
		if !method.IsValid() {
			fmt.Println("Method not found!")
			continue
		}

		args := []reflect.Value{}
		result :=method.Call(args);

		if(result[0].Interface() == nil ) {
			var migrationModel2  model.Migration;
			migrationModel2.Batch = currentBatch;
			migrationModel2.Migration = file.Name()
			config.DB.Create(&migrationModel2)
		}
	}
}
