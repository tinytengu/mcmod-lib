package main

import (
	"fmt"
	"mcmodlib/cmds"
	"mcmodlib/models"
	"os"
	"strings"

	"github.com/tinytengu/go-argparse"
)

var commands = []models.Command{
	cmds.HandlerCmd,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: mcmod <command> <args>")

		fmt.Println("\nCommands:")
		for _, cmd := range commands {
			fmt.Printf("\t%v\n", cmd.SprintInline())
		}
		return
	}

	for _, cmd := range commands {
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
