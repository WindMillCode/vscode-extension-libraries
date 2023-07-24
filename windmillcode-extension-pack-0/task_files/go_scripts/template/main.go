package main

import (
	"github.com/WindMillCode/vscode-extension-libraries/tree/main/windmillcode-extension-pack-0/task_files/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToTestNGApp()

	utils.RunCommand("", []string{})
}
