package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/bopher/bopher/internal/helpers"
	"github.com/bopher/utils"
	"github.com/spf13/cobra"
)

// NewCommand create new app
var NewCommand = new(cobra.Command)

func init() {
	NewCommand.Use = "new [app name]"
	NewCommand.Short = "create new bohper app"
	NewCommand.Args = cobra.MinimumNArgs(1)
	NewCommand.Run = func(cmd *cobra.Command, args []string) {
		// name
		name := args[0]
		root := "./" + name

		if exists, _ := utils.FileExists(root); exists {
			helpers.Handle(fmt.Sprintf("directory %s exists!\n", name))
		}

		// download and extract template
		func(basePath string) {
			dest := basePath + "/bopher.zip"
			helpers.Handle(helpers.Download("https://github.com/bopher/boilerplate/archive/latest.zip", dest))
			helpers.Handle(helpers.Unzip(dest, path.Join("./", name)))
			helpers.Handle(os.Remove(dest))
		}(root)

		// Run wizard
		setup(name, runWizard())
		fmt.Printf("\n\n")

		// Clean template files
		func(basePath string) {
			files := utils.FindFile(basePath, ".tpl.*")
			for _, f := range files {
				os.Remove("./" + f)
			}
		}(root)

		// Tidy app
		func(basePath string) {
			cmd := exec.Command("go", "mod", "tidy")
			cmd.Dir = basePath
			fmt.Printf("tidy app: ")
			if err := cmd.Run(); err != nil {
				fmt.Printf("FAILED!\n")
				fmt.Println(err)
			} else {
				fmt.Printf("OK!\n")
			}
		}(root)

		// Format app
		func(basePath string) {
			cmd := exec.Command("go", "fmt", "./...")
			cmd.Dir = basePath
			fmt.Printf("formatting: ")
			if err := cmd.Run(); err != nil {
				fmt.Printf("FAILED!\n")
				fmt.Println(err)
			} else {
				fmt.Printf("OK!\n")
			}
		}(root)

		// Final message
		fmt.Printf("\nApp created.\nEnjoy it!\n\n")
	}
}
