package command

import (
	"fmt"
	"gogo/utils"
	"os"
	"time"
)

func Make() {
	make := os.Args[2]
	switch make {
	case "model":
		MakeModel()
	}

}

func MakeModel() {
	name := os.Args[3]
	fileName := utils.ToSnakeCase(name)
	dest := "model/" + fileName + ".go"
	modelName := utils.ToPascelCase(name)
	moduleName, _ := utils.GetModuleName()

	utils.CopyFile("copy/model/gogo.go", dest, moduleName, "GOGO", modelName)
	MakeMigrationForModel(name)
}

func MakeMigrationForModel(name string) {
	fmt.Println("fefef")
	fmt.Println(os.Args[3])
	if len(os.Args) > 4 && os.Args[4] == "--migration" {
		currentTime := time.Now()

		// Format the date as "YYYY_MM_DD_HHMMSS"
		timeString := currentTime.Format("2006_01_02_150405")

		fileName := timeString + "_create_" + utils.ToSnakeCase(name) + "_table.go"
		dest := "migration/" + fileName
		moduleName, _ := utils.GetModuleName()
		modelName := utils.ToPascelCase(name)
		fmt.Println(moduleName)
		utils.CopyFile("copy/migration/test_migration.go", dest, moduleName, "GOGO", modelName)
	}
}
