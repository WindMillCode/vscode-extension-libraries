package main

import (
	"go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	sourceBranch := "dev"
	currentBranch := utils.RunCommandAndGetOutput("git",[]string{"rev-parse","--abbrev-ref","HEAD"})
	deleteBranch := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"the local branch to delete:"},
			Default: currentBranch,
		},
	)
	createBranch := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"the local branch to create:"},
			Default: currentBranch,
		},
	)

	utils.RunCommand("git",[]string{"checkout",sourceBranch})
	utils.RunCommand("git",[]string{"pull","origin",sourceBranch})
	utils.RunCommand("git",[]string{"branch","-D",deleteBranch})
	utils.RunCommand("git",[]string{"checkout",sourceBranch})

}

