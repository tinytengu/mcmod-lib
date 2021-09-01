package main

import (
	"fmt"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/commands"
	"os"
	"strings"
)

// Console commands list
var commandsList = []command.Command{
	commands.InitCommand,
	commands.InstallCommand,
	commands.UninstallCommand,
	commands.ListCommand,
}

func main() {
	// No commands or --help
	if len(os.Args) == 1 || len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Println("Usage: mcmod <command> <args>")

		fmt.Println("\nCommands:")
		for _, cmd := range commandsList {
			fmt.Printf("\t%v", cmd.SprintInline())
		}
		return
	}

	// Command passed
	for _, cmd := range commandsList {
		if cmd.Name == os.Args[1] {
			result, err := cmd.Parameters.Parse(strings.Join(os.Args[2:], " "))
			if err != nil {
				fmt.Printf("* %v\n\n", err)
				cmd.Print()
				return
			}
			cmd.Handler(cmd, result)
			return
		}
	}
}
