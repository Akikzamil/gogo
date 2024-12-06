package command

import "gogo/utils"

func RunMigration() {
	utils.RunCommand("go", "run", "main.go", "migrate")
}