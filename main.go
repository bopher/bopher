package main

import (
	"github.com/bopher/bopher/commands"
	"github.com/bopher/cli"
)

func main() {
	cli := cli.NewCLI("bopher", "Bopher framework cli tools")
	cli.AddCommand(commands.VersionCommand)
	cli.AddCommand(commands.NewCommand)
	cli.Run()
}
