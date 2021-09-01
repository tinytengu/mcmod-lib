package commands

import (
	"fmt"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/environment"
	"mcmodlib/pkg/utils"
	"path/filepath"

	"github.com/tinytengu/go-argparse"
	"github.com/tinytengu/go-cfapi"
)

// 'install' command
var InstallCommand = command.Command{
	Name: "install",
	Desc: "Install a mod or mods",
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
		envPath, _ := filepath.Abs(args.Flags["env"])
		env := environment.NewEnvironment(envPath)

		// Check if environment already exists
		if !utils.IsFileExists(env.GetConfigPath()) {
			fmt.Printf("Modding environment in folder '%v' is not found\n", envPath)
			return
		}

		api := cfapi.NewApi()

		// Cycle through all passed mods
		for _, selector := range args.Args {
			// Get mod from CurseForge
			mod, err := api.GetMod(selector)
			if err != nil {
				fmt.Printf("Error (%v): %v", selector, err)
				continue
			}
			fmt.Printf("Installing %v\n", mod.Title)

			// Download .jar file directly to the environment folder
			cdn := mod.Download.GetCDN()
			err = utils.DownloadFile(cdn, filepath.Join(envPath, mod.Download.Display))
			if err != nil {
				fmt.Printf("Error (%v): %v\n", selector, err)
				continue
			}

			// Add mod to the environment file
			env.Mods = append(env.Mods, environment.Mod{
				Id:    selector,
				Mcver: mod.Download.Version,
				Type:  mod.Download.Type,
				File:  mod.Download.Display,
			})
		}

		// Write out changes
		env.Write()
		fmt.Println("All mods are successfully installed")
	},
}
