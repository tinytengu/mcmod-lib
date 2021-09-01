package commands

import (
	"fmt"
	"log"
	"mcmodlib/pkg/command"
	"mcmodlib/pkg/environment"
	"mcmodlib/pkg/utils"
	"os"
	"path/filepath"
	"strings"

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
				Optional: true,
			},
		},
		Flags: argparse.FlagsList{
			"r": {
				Desc:     "Mods list reference file",
				Default:  "",
				Optional: true,
			},
			"env": {
				Desc:     "Environment path",
				Default:  ".",
				Optional: true,
			},
		},
	},
	Handler: func(cmd command.Command, args argparse.ParseResult) {
		// REFACTOR: Make it prettier pls
		if len(args.Flags["r"]) == 0 && len(args.Args[0]) == 0 {
			keys := make([]string, 0, len(cmd.Parameters.Args))
			for k := range cmd.Parameters.Args {
				keys = append(keys, k)
			}

			fmt.Printf("* `%v` argument required\n", keys[0])
			return
		}

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
		var selectors []string

		refFile := args.Flags["r"]
		if len(refFile) != 0 {
			if !utils.IsFileExists(refFile) {
				fmt.Printf("File '%v' not found\n", refFile)
				return
			}
			data, err := utils.ReadFile(refFile)
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = utils.IsValidModsList(string(data))
			if err != nil {
				log.Fatal(err)
				return
			}
			for _, v := range strings.Split(string(data), "\n") {
				if len(strings.TrimSpace(v)) == 0 {
					continue
				}
				selectors = append(selectors, v)
			}
		} else {
			selectors = args.Args
		}

		// Cycle through all passed mods
		for _, selector := range selectors {
			sel := utils.ParseSelector(selector)
			if len(sel.Id) == 0 {
				fmt.Printf("Error: invalid mod id '%v'", sel.Id)
				continue
			}
			if len(env.Properties["mcver"]) != 0 && len(sel.Version) == 0 {
				sel.Version = env.Properties["mcver"]
			}
			if len(env.Properties["type"]) != 0 && len(sel.Type) == 0 {
				sel.Type = env.Properties["type"]
			}

			// Get mod from CurseForge
			mod, err := api.GetMod(sel.Id)
			if err != nil {
				fmt.Printf("Error (%v): %v\n", sel.Id, err)
				continue
			}
			modId := utils.SliceLast(strings.Split(mod.Urls.Curseforge, "/"))

			// Filter mod files
			files := utils.FilterModFiles(mod, sel)
			if len(files) == 0 {
				fmt.Printf("Error (%v): no matching files found\n", modId)
				continue
			}
			file := files[0]

			fmt.Printf("Installing %v\n", mod.Title)

			// Delete old .jar file
			modIdx, envMod := env.Mods.GetById(sel.Id)
			if modIdx != -1 {
				os.Remove(filepath.Join(envPath, envMod.File))
			}

			// Download .jar file directly to the environment folder
			cdn := file.GetCDN()
			err = utils.DownloadFile(cdn, filepath.Join(envPath, file.Display))
			if err != nil {
				fmt.Printf("Error (%v): %v\n", modId, err)
				continue
			}

			if modIdx == -1 {
				// Add mod to the environment
				env.Mods = append(env.Mods, environment.Mod{
					Id:    modId,
					Mcver: file.Version,
					Type:  file.Type,
					File:  file.Display,
				})
			} else {
				// Update mod environment properties
				env.Mods[modIdx].Id = modId
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
