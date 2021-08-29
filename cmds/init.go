package cmds

import (
	"fmt"
	"mcmodlib/models"
	"mcmodlib/shared"
	"path/filepath"

	"github.com/tinytengu/go-argparse"
)

var HandlerCmd = models.Command{
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

func handler(cmd models.Command, args argparse.ParseResult) {
	// Determine initialization path (cwd or passed)
	initPath, _ := filepath.Abs(args.Args[0])
	// var initPath string
	// if len(args.Args) == 0 {
	// 	initPath, _ = filepath.Abs(".")
	// } else {
	// 	initPath = args.Args[0]
	// }

	env := models.NewEnvironment(initPath)

	if shared.IsFileExists(env.GetConfigPath()) {
		fmt.Printf("Modding environment in folder '%v' already exists\n", initPath)
		return
	}

	env.Write()
	fmt.Printf("Modding environment initialized in '%v' folder\n", initPath)
}
