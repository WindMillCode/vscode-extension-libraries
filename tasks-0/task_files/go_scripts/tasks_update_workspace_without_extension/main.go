package main

import (
	"encoding/json"
	"fmt"
	"os"
	"main/shared"
	"github.com/windmillcode/go_cli_scripts/v3/utils"
)



func main() {
	utils.CDToWorkspaceRoot()
	workSpaceFolder, _ := os.Getwd()
	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the executable to use (try with windmillcode_go first if not then use go)",
		Choices: []string{"go", "windmillcode_go"},
		Default: "windmillcode_go",
	}
	goExecutable := utils.ShowMenu(cliInfo, nil)

	proceed := "FALSE"

	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, ".vscode/tasks.json")

	content, err, shouldReturn := createTasksJson(tasksJsonFilePath, false)
	if shouldReturn {
		return
	}

	var tasksJSON shared.TasksJSON
	err = json.Unmarshal(content, &tasksJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	goScriptsDestDirPath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, "ignore/Windmillcode/go_scripts")

	shared.RebuildExecutables(proceed, cliInfo, tasksJSON, goScriptsDestDirPath, goExecutable)

}



func createTasksJson(tasksJsonFilePath string, triedCreateOnError bool) ([]byte, error, bool) {
	content, err := os.ReadFile(tasksJsonFilePath)
	if err != nil {
		if triedCreateOnError {
			return nil, err, true
		}

		// If the file doesn't exist, create it.
		_, createErr := os.Create(tasksJsonFilePath)
		if createErr != nil {
			return nil, createErr, true
		}

		// Recursively attempt to read the file after creating it.
		return createTasksJson(tasksJsonFilePath, true)
	}

	return content, nil, false
}


