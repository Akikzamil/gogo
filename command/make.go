package command

import (
	"gogo/utils"
	"os"
	"time"
)

func Make() {
	make := os.Args[2]
	switch make {
	case "model":
		MakeModel()
	case "migration":
		makeMigration()
	}

}

func MakeModel() {
	name := os.Args[3]
	fileName := utils.ToSnakeCase(name)
	dest := "model/" + fileName + ".go"
	modelName := utils.ToPascelCase(name)
	moduleName, _ := utils.GetModuleName()

	utils.CopyFile("copy/model/gogo.go", dest, moduleName, "GOGO", modelName,"","","","")
	MakeMigrationForModel(name)
}

func MakeMigrationForModel(name string) {
	if len(os.Args) > 4 && os.Args[4] == "--migration" {
		currentTime := time.Now()

		// Format the date as "YYYY_MM_DD_HHMMSS"
		timeString := currentTime.Format("2006_01_02_150405")

		fileName := timeString + "_create_" + utils.ToSnakeCase(name) + "_table.go"
		fileNameWithOutExt := timeString + "_create_" + utils.ToSnakeCase(name) + "_table"
		dest := "migration/" + fileName
		moduleName, _ := utils.GetModuleName()
		modelName := utils.ToPascelCase(name)

		utils.CopyFile("copy/migration/test_migration.go", dest, moduleName, "GOGO", modelName,"Up","Up"+fileNameWithOutExt, "Down", "Down"+fileNameWithOutExt)
	}
}

func makeMigration() {
	if len(os.Args)>3{

		currentTime := time.Now()

		// Format the date as "YYYY_MM_DD_HHMMSS"
		timeString := currentTime.Format("2006_01_02_150405")
		name :=utils.ToSnakeCase(os.Args[3]);
		moduleName, _ := utils.GetModuleName();
		fileNameWithoutExt:= timeString + "_modify_" +  name + "_table"
		fileName := fileNameWithoutExt + ".go"
		dest := "migration/" + fileName

		utils.CopyFile("copy/migration/modify_migration.go", dest, moduleName, "Up2","Up"+fileNameWithoutExt, "Down2", "Down"+fileNameWithoutExt ,"","")

	}

}
