package commands

import (
	"fmt"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/environment"
	"mcmodlib/pkg/utils"
	"path/filepath"

	"github.com/tinytengu/go-argparse"
)

// 'init' command
var InitCommand = command.Command{
	Name: "init",
	Desc: "Initialize modding environment",
	Args: argparse.ArgsList{
		"path": {
			Desc:     "Initialization path",
			Default:  ".",
			Optional: false,
			Flag:     false,
		},
	},
	Handler: handler,
}

// Command handler
func handler(cmd command.Command, args argparse.ParseResult) {
	initPath, _ := filepath.Abs(args.Args[0])
	env := environment.NewEnvironment(initPath)

	if utils.IsFileExists(env.GetConfigPath()) {
		fmt.Printf("Modding environment in folder '%v' already exists\n", initPath)
		return
	}

	env.Write()
	fmt.Printf("Modding environment initialized in '%v' folder\n", initPath)
}
