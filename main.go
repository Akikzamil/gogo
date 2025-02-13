package main

import (
	"fmt"
	"gogo/command"
	"os"
)

func main() {
	if(len(os.Args)==1) {
		fmt.Println("Wellcome to golang");
		return;
	}
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
