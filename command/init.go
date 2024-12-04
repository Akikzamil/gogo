package command

import (
	"fmt"
	"gogo/utils"
	"os"
)

var ProjectName = "gogo"

func Init() {
	if !moduleInit() {
		return
	}

	if !insallGoFiber() {
		fmt.Println("fiber is not installed")
		return
	}

	if !installEnvLibrary() {
		return
	}

	if !installGorm() {
		return
	}

	if !copyMainFile() {
		return
	}

	copyOtherFile()

}

func moduleInit() bool {
	if len(os.Args) > 2 {
		ProjectName = os.Args[2]
	}

	isInstalled := utils.RunCommand("go", "mod", "init", ProjectName)

	if isInstalled {
		fmt.Println("go module successfully installed")
	}
	return isInstalled
}

func copyMainFile() bool {
	return utils.CopyFile("copy/main.go", "main.go", ProjectName)
}

func insallGoFiber() bool {
	isInstalled := utils.RunCommand("go", "get", "github.com/gofiber/fiber/v2")
	if isInstalled {
		fmt.Println("go fiber successfully installed")
	}

	return isInstalled
}

func installEnvLibrary() bool {
	isInstalled := utils.RunCommand("go", "get", "github.com/joho/godotenv")
	if isInstalled {
		fmt.Println("go env successfully installed")
	}

	isEnvFileCopied := utils.CopyFile("copy/.env.example", "env.example", ProjectName)

	if isInstalled {
		fmt.Println("env File copied succesfully")
	}

	return isInstalled && isEnvFileCopied
}

func installGorm() bool {
	isGormInstalled := utils.RunCommand("go", "get", "-u", "gorm.io/gorm")
	if isGormInstalled {
		fmt.Println("gorm successfully installed")
	}

	isSqliteDriverInstalled := utils.RunCommand("go", "get", "gorm.io/driver/sqlite")

	isMysqlnstalled := utils.RunCommand("go", "get", "gorm.io/driver/mysql")

	isPosgresSqlInstalled := utils.RunCommand("go", "get", "gorm.io/driver/postgres")

	return isGormInstalled && isMysqlnstalled && isPosgresSqlInstalled && isSqliteDriverInstalled
}

type Filepath struct {
	copyPath  string
	pastePath string
}

var allfilePaths []Filepath = []Filepath{
	{copyPath: "copy/config/database.go", pastePath: "config/database.go"},
	{copyPath: "copy/utils/env.go", pastePath: "utils/env.go"},

}

func copyOtherFile() {
	for _, v := range allfilePaths {
		utils.CopyFile(v.copyPath, v.pastePath, ProjectName)
	}
}
