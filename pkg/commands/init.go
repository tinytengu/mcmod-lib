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
		"mcver": {
			Desc:     "Default Minecraft version",
			Default:  "",
			Optional: true,
			Flag:     true,
		},
		"modtype": {
			Desc:     "Default mod type (alpha, beta, release)",
			Default:  "",
			Optional: true,
			Flag:     true,
		},
	},
	Handler: handler,
}

// Command handler
func handler(cmd command.Command, args argparse.ParseResult) {
	initPath, _ := filepath.Abs(args.Args[0])
	env := environment.NewEnvironment(initPath)

	// Check if environment already exists
	if utils.IsFileExists(env.GetConfigPath()) {
		fmt.Printf("Modding environment in folder '%v' already exists\n", initPath)
		return
	}

	// Validate 'mcver' flag
	env.Storage.Properties.ValidateFlag(
		args.Flags, "mcver",
		utils.IsValidMcVersion,
		fmt.Sprintf("Invalid Minecraft version: %v\n", args.Flags["mcver"]),
	)

	// Validate 'modtype' flag
	env.Storage.Properties.ValidateFlag(
		args.Flags, "modtype",
		utils.IsValidModType,
		fmt.Sprintf("Invalid mod type: %v\n", args.Flags["modtype"]),
	)

	env.Write()
	fmt.Printf("Modding environment initialized at '%v' folder\n", initPath)
}
