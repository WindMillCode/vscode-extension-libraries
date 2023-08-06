package main

import (
	"path/filepath"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	cliInfo := utils.ShowMenuModel{
		Prompt: "choose the package manager",
		Choices:[]string{"npm","yarn"},
		Default:"npm",
	}
	packageManager := utils.ShowMenu(cliInfo,nil)
	cliInfo = utils.ShowMenuModel{
		Other: true,
		Prompt:  "Choose the node.js app",
		Choices: []string{filepath.Join("./apps/frontend/AngularApp"), filepath.Join(".\\apps\\cloud\\FirebaseApp")},
	}
	appLocation := utils.ShowMenu(cliInfo, nil)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)
	utils.CDToLocation(appLocation)
	if reinstall == "true" {
		if packageManager == "npm" {
			utils.RunCommand("rm", []string{"package-lock.json"})
		} else {
			utils.RunCommand("rm", []string{"yarn.lock"})
		}
		utils.RunCommand("rm", []string{"-r", "node_modules"})
		utils.RunCommand(packageManager, []string{"cache", "clean"})

	}
	if packageManager == "npm" {
		utils.RunCommand(packageManager, []string{"install","-s"})
	} else{
		utils.RunCommand(packageManager, []string{"install"})
	}
}
