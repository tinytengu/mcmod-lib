package command

import (
	"fmt"

	"github.com/tinytengu/go-argparse"
)

type Command struct {
	Name    string
	Desc    string
	Args    argparse.ArgsList
	Handler CommandHandler
}

func (cmd *Command) Print() {
	fmt.Printf("%v - %v\n\n", cmd.Name, cmd.Desc)
	argsSet := (argparse.ArgsSet{
		Args: cmd.Args,
	})
	argsSet.Print()
}

func (cmd *Command) PrintInline() {
	fmt.Printf("%v - %v\n\n", cmd.Name, cmd.Desc)
}

func (cmd *Command) SprintInline() string {
	return fmt.Sprintf("%v - %v\n\n", cmd.Name, cmd.Desc)
}
