package commands

import (
	"fmt"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/environment"
	"mcmodlib/pkg/utils"
	"os"
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
		// Initialize environment structure
		envPath, _ := filepath.Abs(args.Flags["env"])
		env := environment.NewEnvironment(envPath)
		env.Read()

		// Check if environment already exists
		if !utils.IsFileExists(env.GetConfigPath()) {
			fmt.Printf("Modding environment in folder '%v' is not found\n", envPath)
			return
		}

		// Initialize CurseForge API
		api := cfapi.NewApi()
		installed := 0

		// Cycle through all passed mods
		for _, selector := range args.Args {
			sel := utils.ParseSelector(selector)
			if len(sel.Id) == 0 {
				fmt.Printf("Error: invalid mod id '%v'", sel.Id)
				continue
			}
			if len(env.Properties["mcver"]) != 0 && len(sel.Version) == 0 {
				sel.Version = env.Properties["mcver"]
			}
			if len(env.Properties["modtype"]) != 0 && len(sel.Type) == 0 {
				sel.Type = env.Properties["modtype"]
			}

			// Get mod from CurseForge
			mod, err := api.GetMod(sel.Id)
			if err != nil {
				fmt.Printf("Error (%v): %v\n", sel.Id, err)
				continue
			}

			// Filter mod files
			files := utils.FilterModFiles(mod, sel)
			if len(files) == 0 {
				fmt.Printf("Error (%v): no matching files found\n", sel.Id)
				continue
			}
			file := files[0]

			fmt.Printf("Installing %v\n", mod.Title)

			// Delete old .jar file
			modIdx, envMod := env.Mods.GetById(sel.Id)
			os.Remove(filepath.Join(envPath, envMod.File))

			// Download .jar file directly to the environment folder
			cdn := file.GetCDN()
			err = utils.DownloadFile(cdn, filepath.Join(envPath, file.Display))
			if err != nil {
				fmt.Printf("Error (%v): %v\n", sel.Id, err)
				continue
			}

			if modIdx == -1 {
				// Add mod to the environment
				env.Mods = append(env.Mods, environment.Mod{
					Id:    sel.Id,
					Mcver: file.Version,
					Type:  file.Type,
					File:  file.Display,
				})
			} else {
				// Update mod environment properties
				env.Mods[modIdx].Id = sel.Id
				env.Mods[modIdx].Mcver = file.Version
				env.Mods[modIdx].Type = file.Type
				env.Mods[modIdx].File = file.Display
			}
			installed++
		}

		// Write out changes
		env.Write()
		if installed == 0 {
			fmt.Println("Unable to install required mods")
		} else {
			fmt.Printf("Mods are successfully installed (%v)\n", installed)
		}
	},
}
