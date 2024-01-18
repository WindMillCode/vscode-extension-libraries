package main

import (
	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	cliInfo := utils.ShowMenuModel{
		Prompt: "choose the package manager",
		Choices: []string{"npm","yarn","pnpm"},
		Default :"npm",
	}
	exectuable := utils.ShowMenu(cliInfo,nil)
	cliInfo = utils.ShowMenuModel{
		Prompt:  "Choose an option:",
		Choices: []string{"dev", "preview", "prod"},
	}
	envType := utils.ShowMenu(cliInfo, nil)
	utils.CDToAngularApp()
	if exectuable == "npm"{
		utils.RunCommand(exectuable, []string{"run","analyze:" + envType})
	} else if exectuable == "pnpm"{
		utils.RunCommand(exectuable, []string{"run","analyze:" + envType})
	} else if exectuable == "yarn"{
		utils.RunCommand(exectuable, []string{"analyze:" + envType})
	}
}
