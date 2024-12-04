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
	}
}
