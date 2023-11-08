package main

import (
	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	cliInfo := utils.ShowMenuModel{
		Prompt:  "Choose an option:",
		Choices: []string{"dev", "preview", "prod"},
	}
	envType := utils.ShowMenu(cliInfo, nil)
	utils.CDToAngularApp()
	utils.RunCommand("yarn", []string{"analyze:" + envType})
}
