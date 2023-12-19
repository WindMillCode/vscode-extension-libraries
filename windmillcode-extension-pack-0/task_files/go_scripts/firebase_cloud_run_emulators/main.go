package main

import (
	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToFirebaseApp()

	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the package manager",
		Choices: []string{"npm", "yarn"},
		Default: "npm",
	}
	packageManager := utils.ShowMenu(cliInfo, nil)

	utils.RunCommand(packageManager, []string{"run","cleanup"})
	utils.RunCommand("npx", []string{"firebase", "emulators:start", "--import=devData", "--export-on-exit"})
}
