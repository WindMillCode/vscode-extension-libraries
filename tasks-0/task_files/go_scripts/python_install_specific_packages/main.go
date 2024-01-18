package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceFolder, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	settings, err := utils.GetSettingsJSON(workspaceFolder)
	if err != nil {
		return
	}
	cliInfo := utils.ShowMenuModel{
		Other:  true,
		Prompt: "Choose an option:",
		Choices: []string{
			utils.JoinAndConvertPathToOSFormat("./apps/backend/FlaskApp"),
			utils.JoinAndConvertPathToOSFormat("."),
		},
	}
	appLocation := utils.ShowMenu(cliInfo, nil)
	appLocation = utils.JoinAndConvertPathToOSFormat(appLocation)

	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide a python version for pyenv to use"},
			Default: settings.ExtensionPack.PythonVersion0,
		},
	)
	utils.RunCommand("pyenv", []string{"global", pythonVersion})

	packageList := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt: "Provide the names of the packages you would like to install",
			ErrMsg: "You must provide packages for installation",
		},
	)

	cliInfo = utils.ShowMenuModel{
		Prompt:  "reinstall?",
		Choices: []string{"true", "false"},
	}
	reinstall := utils.ShowMenu(cliInfo, nil)
	utils.CDToLocation(appLocation)
	var sitePackages string
	targetOs := runtime.GOOS
	switch targetOs {
	case "windows":

		sitePackages = utils.JoinAndConvertPathToOSFormat(".", "site-packages", "windows")

	case "linux", "darwin":
		sitePackages = utils.JoinAndConvertPathToOSFormat(".", "site-packages", "linux")

	default:
		fmt.Println("Unknown Operating System:", targetOs)
	}
	if reinstall == "true" {
		utils.RunCommand("pip", []string{"uninstall", packageList.InputString})
	}
	utils.RunCommand("pip", []string{"install", packageList.InputString, "--target", sitePackages})
	utils.RunCommand("pip", []string{"freeze", "--all", "--path", sitePackages})

}
