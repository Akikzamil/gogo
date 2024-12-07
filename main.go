package main

import (
	"gogo/command"
	"os"
)

func main() {
	commandName := os.Args[1];

	switch(commandName) {
	case "init":
		command.Init();
	case "make":
		command.Make();
	case "migrate:rollback":
		command.RunDownMigration();
	case "migrate":
		command.RunMigration();
	}
}
