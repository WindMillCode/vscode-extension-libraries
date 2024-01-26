package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {
	scriptFolder,err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the script dir")
	}
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
	appLocation = utils.JoinAndConvertPathToOSFormat(workspaceFolder,appLocation)

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
		Prompt:  "uninstall?",
		Choices: []string{"true", "false"},
	}
	uninstall := utils.ShowMenu(cliInfo, nil)
	cliInfo = utils.ShowMenuModel{
		Prompt:  "install?",
		Choices: []string{"true", "false"},
	}
	install := utils.ShowMenu(cliInfo, nil)
	utils.CDToLocation(appLocation)
	var sitePackages string
	targetOs := runtime.GOOS

	cliInfoMulti := utils.ShowMenuMultipleModel{
		Prompt: "Select the requirements files ",
		Choices:[]string{"windows-requirements.txt","linux-requirements.txt","darwin-requirements.txt"},
		SelectionLimit: 3,
		Defaults: []string{"windows-requirements.txt","linux-requirements.txt","darwin-requirements.txt"},
	}
	requirementsFiles := utils.ShowMenuMultipleOptions(cliInfoMulti,nil)
	switch targetOs {
	case "windows":

		sitePackages = utils.JoinAndConvertPathToOSFormat(appLocation, "site-packages", "windows")

	case "linux", "darwin":
		sitePackages = utils.JoinAndConvertPathToOSFormat(appLocation, "site-packages", "linux")

	default:
		fmt.Println("Unknown Operating System:", targetOs)
	}
	if uninstall == "true" {

		packageListArgs := append([]string{"remove_local_packages.py"}, packageList.InputArray...)
		packageListArgs = append(packageListArgs, sitePackages)
		removeLocalPagesOptions :=utils.CommandOptions{
			Command:        "python",
			Args:           packageListArgs,
			TargetDir:      scriptFolder,
		}
		utils.RunCommandWithOptions(removeLocalPagesOptions)
		for _, requirementsFile := range requirementsFiles {
			requirementsFilePath := utils.JoinAndConvertPathToOSFormat(appLocation, requirementsFile)
			err := utils.RemoveContentFromFile(requirementsFilePath, packageList.InputArray)
			if err != nil {
					fmt.Printf("Error adding package to requirements file: %v\n", err)
			}
		}

	}
	if install == "true" {
		fmt.Println(packageList.InputArray)
		packageListArgs := append([]string{"install"}, packageList.InputArray...)
		packageListArgs = append(packageListArgs,"--target", sitePackages)
		utils.RunCommand("pip", packageListArgs)
		for _, requirementsFile := range requirementsFiles {
			requirementsFilePath := utils.JoinAndConvertPathToOSFormat(appLocation, requirementsFile)
			for _, packageName := range packageList.InputArray {
				err := utils.AddContentToFile(requirementsFilePath, "\n"+packageName, "suffix")
				if err != nil {
						fmt.Printf("Error adding package to requirements file: %v\n", err)
				}
			}
		}
	}


}
