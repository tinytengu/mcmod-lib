package main

import (
	"fmt"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/commands"
	"os"
	"strings"

	"github.com/tinytengu/go-argparse"
)

var commandsList = []command.Command{
	commands.HandlerCmd,
}

func main() {
	if len(os.Args) == 1 || len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Println("Usage: mcmod <command> <args>")

		fmt.Println("\nCommands:")
		for _, cmd := range commandsList {
			fmt.Printf("\t%v", cmd.SprintInline())
		}
		return
	}

	for _, cmd := range commandsList {
		if cmd.Name == os.Args[1] {
			argsSet := argparse.ArgsSet{
				Args: cmd.Args,
			}
			result, err := argsSet.Parse(strings.Join(os.Args[2:], " "))
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
