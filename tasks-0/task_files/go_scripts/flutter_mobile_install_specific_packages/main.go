package main

import (
	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	cliInfo := utils.ShowMenuModel{
		Other:   true,
		Prompt:  "Choose an option:",
		Choices: []string{".\\apps\\mobile\\FlutterApp"},
	}
	appLocation := utils.ShowMenu(cliInfo, nil)
	appLocation = utils.JoinAndConvertPathToOSFormat(appLocation)

	packageList := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt: "Provide the names of the packages you would like to install",
			ErrMsg: "You must provide packages for installation",
		},
	)

	utils.CDToLocation(appLocation)

	utils.RunCommand("flutter", []string{"pub", "add", packageList.InputString})

}
