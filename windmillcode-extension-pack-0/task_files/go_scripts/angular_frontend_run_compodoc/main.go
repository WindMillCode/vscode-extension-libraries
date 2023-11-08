package main

import (
	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToAngularApp()

	utils.RunCommand("yarn", []string{"compodoc:build-and-serve"})
}
