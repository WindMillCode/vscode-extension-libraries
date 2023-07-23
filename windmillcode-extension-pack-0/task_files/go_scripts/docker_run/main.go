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


	utils.RunCommand("docker",[]string{"start",dockerContainerName})
}

