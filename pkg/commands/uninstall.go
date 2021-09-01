package commands

import (
	"fmt"
	"log"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/environment"
	"mcmodlib/pkg/utils"
	"os"
	"path/filepath"
	"regexp"

	"github.com/tinytengu/go-argparse"
)

// 'uninstall' command
var UninstallCommand = command.Command{
	Name: "uninstall",
	Desc: "Uninstall a mod from modding environment",
	Parameters: argparse.Parameters{
		Args: argparse.ArgsList{
			"mods": {
				Desc:     "Mods list",
				Default:  "",
				Optional: false,
			},
		},
		Flags: argparse.FlagsList{
			"env": {
				Desc:     "Environment path",
				Default:  ".",
				Optional: true,
			},
		},
	},
	Handler: func(cmd command.Command, args argparse.ParseResult) {
		// Initialize environment structure
		envPath, _ := filepath.Abs(args.Flags["env"])
		env := environment.NewEnvironment(envPath)
		env.Read()

		// Check if environment already exists
		if !utils.IsFileExists(env.GetConfigPath()) {
			fmt.Printf("Modding environment in folder '%v' is not found\n", envPath)
			return
		}

		var mods environment.ModsList

		for _, arg := range args.Args {
			expr, err := regexp.Compile(arg)
			if err != nil {
				log.Fatal(err)
				return
			}
			for _, mod := range env.Mods {
				if !expr.MatchString(mod.Id) {
					mods = append(mods, mod)
				} else {
					filePath := filepath.Join(env.GetPath(), mod.File)
					if utils.IsFileExists(filePath) {
						os.Remove(filePath)
					}
					fmt.Printf("'%v' has been uninstalled\n", mod.Id)
				}
			}
		}

		env.Mods = mods
		env.Write()
	},
}
