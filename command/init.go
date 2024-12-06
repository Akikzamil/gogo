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
	installOtherDependencies();

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
	return utils.CopyFile("copy/main.go", "main.go", ProjectName,"","")
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

	isEnvFileCopied := utils.CopyFile("copy/.env.example", ".env.example", ProjectName,"","")

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


func installOtherDependencies() bool{
	isFiberJwtInstalled := utils.RunCommand("go", "get", "-u", "github.com/gofiber/contrib/jwt");
	isJwtInstalled := utils.RunCommand("go", "get", "-u", "github.com/golang-jwt/jwt/v5");
	isTidyCommandRun := utils.RunCommand("go", "mod", "tidy");

	return isFiberJwtInstalled && isJwtInstalled && isTidyCommandRun
}

type Filepath struct {
	copyPath  string
	pastePath string
}

var allfilePaths []Filepath = []Filepath{
	{copyPath: "copy/config/database.go", pastePath: "config/database.go"},
	{copyPath: "copy/utils/env.go", pastePath: "utils/env.go"},
	{copyPath: "copy/middleware/auth.go", pastePath: "middleware/auth.go"},
	{copyPath: "copy/route/route.go", pastePath: "route/route.go"},
	{copyPath: "copy/model/migration.go", pastePath: "model/migration.go"},
}

func copyOtherFile() {
	for _, v := range allfilePaths {
		utils.CopyFile(v.copyPath, v.pastePath, ProjectName,"","")
	}
}
