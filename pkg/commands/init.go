// 'init' command
package commands

import (
	"fmt"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/environment"
	"mcmodlib/pkg/utils"
	"path/filepath"

	"github.com/tinytengu/go-argparse"
)

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

	if utils.IsFileExists(env.GetConfigPath()) {
		fmt.Printf("Modding environment in folder '%v' already exists\n", initPath)
		return
	}

	mcVer := args.Flags["mcver"]
	if mcVer != "" {
		if !utils.IsValidMcVersion(mcVer) {
			fmt.Printf("Invalid Minecraft version: %v\n", mcVer)
		} else {
			env.Storage.Properties["mcver"] = mcVer
		}
	}

	modType := args.Flags["modtype"]
	if modType != "" {
		if !utils.IsValidModType(modType) {
			fmt.Printf("Invalid mod type: %v\n", modType)
		} else {
			env.Storage.Properties["modtype"] = modType
		}
	}

	env.Write()
	fmt.Printf("Modding environment initialized in '%v' folder\n", initPath)
}
