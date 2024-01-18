package main

import (
	"encoding/json"
	"fmt"
	"os"
	"main/shared"
	"github.com/windmillcode/go_cli_scripts/v4/utils"
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



	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, ".vscode/tasks.json")

	content, err, shouldReturn := shared.CreateTasksJson(tasksJsonFilePath, false)
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

	shared.RebuildExecutables("FALSE", cliInfo, tasksJSON, goScriptsDestDirPath, goExecutable)

}





