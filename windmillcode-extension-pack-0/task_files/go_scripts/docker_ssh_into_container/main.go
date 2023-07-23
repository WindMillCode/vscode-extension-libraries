package main

import (
	"go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	dockerContainerName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"the name of the container"},
			ErrMsg: "you must provide a container to run",
		},
	)
	cliInfo := utils.ShowMenuModel{
		Prompt: "the command line shell (if uncertain select bash)",
		Choices:[]string{"sh", "bash", "dash", "zsh", "cmd", "fish", "ksh", "powershell"},
	}
	shell := utils.ShowMenu(cliInfo,nil)


	utils.RunCommand("docker",[]string{"exec","-it",dockerContainerName,shell})
}

