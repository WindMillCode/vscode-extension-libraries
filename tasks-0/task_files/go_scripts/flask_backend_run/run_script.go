package main

import (
	"os"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)


func main() {
	utils.CDToWorkspaceRoot()
	utils.CDToFlaskApp()
	flaskAppFolder, err := os.Getwd()
	if err != nil {
		return
	}
	flaskAppCommand := utils.CommandOptions{
		Command:            "python",
		Args:               []string{"app.py", "--use_reloader", "False"},
		GetOutput:          false,
		PrintOutputOnly:    true,
		NonBlocking:        true,
		TargetDir:          flaskAppFolder,
	}
	utils.RunCommandWithOptions(flaskAppCommand)
}
