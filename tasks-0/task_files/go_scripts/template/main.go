package main

import (
	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToTestNGApp()

	opts := utils.CommandOptions{
		Command: "",
		Args: []string{},
	}
	utils.RunCommandWithOptions(opts)
}
