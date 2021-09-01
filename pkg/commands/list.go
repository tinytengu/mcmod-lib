package commands

import (
	"fmt"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/environment"
	"mcmodlib/pkg/utils"
	"path/filepath"

	"github.com/tinytengu/go-argparse"
)

// 'list' command
var ListCommand = command.Command{
	Name: "list",
	Desc: "List of installed mods",
	Parameters: argparse.Parameters{
		Args: argparse.ArgsList{
			"path": {
				Desc:     "Envronment path",
				Default:  ".",
				Optional: true,
			},
		},
		Flags: argparse.FlagsList{},
	},
	Handler: func(cmd command.Command, args argparse.ParseResult) {
		envPath, _ := filepath.Abs(args.Args[0])
		env := environment.NewEnvironment(envPath)
		env.Read()

		// Check if environment already exists
		if !utils.IsFileExists(env.GetConfigPath()) {
			fmt.Printf("Modding environment in folder '%v' is not found\n", envPath)
			return
		}

		for _, mod := range env.Mods {
			fmt.Printf("%v:%v:%v:%v\n", mod.Id, mod.Mcver, mod.Type, mod.File)
		}
	},
}
