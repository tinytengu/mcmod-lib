package command

import (
	"fmt"

	"github.com/tinytengu/go-argparse"
)

type Command struct {
	Name       string
	Desc       string
	Parameters argparse.Parameters
	Handler    CommandHandler
}

func (cmd *Command) Print() {
	fmt.Printf("%v - %v\n\n", cmd.Name, cmd.Desc)
	cmd.Parameters.Print()
}

func (cmd *Command) PrintInline() {
	fmt.Printf("%v - %v\n", cmd.Name, cmd.Desc)
}

func (cmd *Command) SprintInline() string {
	return fmt.Sprintf("%v - %v\n\n", cmd.Name, cmd.Desc)
}
