package main

import (
	"encoding/json"
	"fmt"
	"main/shared"
	"os"
	"regexp"
	"strings"

	"github.com/windmillcode/go_cli_scripts/v3/utils"
)



func main() {
	workSpaceFolder := os.Args[1]
	extensionFolder := os.Args[2]
	tasksJsonRelativeFilePath := os.Args[3]
	cliInfo := utils.ShowMenuModel{
		Prompt:  "choose the executable to use (try with windmillcode_go first if not then use go)",
		Choices: []string{"go", "windmillcode_go"},
		Default: os.Args[4],
	}
	goExecutable := utils.ShowMenu(cliInfo, nil)
	cliInfo = utils.ShowMenuModel{
		Prompt:  "This will delete your vscode/tasks.json in your workspace folder. If you don't have a .vscode/tasks.json or you have not used this command before, it is safe to choose TRUE. Otherwise, consult with your manager before continuing",
		Choices: []string{"TRUE", "FALSE"},
	}
	proceed := utils.ShowMenu(cliInfo, nil)

	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(extensionFolder, tasksJsonRelativeFilePath)

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
	goScriptsSourceDirPath := utils.JoinAndConvertPathToOSFormat(extensionFolder, "task_files/go_scripts")
	goScriptsDestDirPath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, "ignore/Windmillcode/go_scripts")

	if proceed == "TRUE" {

		for index, task := range tasksJSON.Tasks {

			pattern0 := ":"
			regex0 := regexp.MustCompile(pattern0)
			programLocation0 := regex0.Split(task.Label, -1)
			pattern1 := " "
			regex1 := regexp.MustCompile(pattern1)
			programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
			programLocation2 := strings.Join(programLocation1, "_")
			programLocation3 := "ignore//${input:current_user_0}//go_scripts//" + programLocation2
			linuxTaskExecutable := ".//main"
			if task.Label == "tasks: update workspace without extension" {
				linuxTaskExecutable = "go run main.go"
			}
			linuxCommand0 := "cd " + programLocation3 + " ; " + linuxTaskExecutable
			windowsCommand0 := "cd " + strings.Replace(programLocation3, "//", "\\", -1) + " ; " + strings.Replace(linuxTaskExecutable, "//", "\\", -1)
			tasksJSON.Tasks[index].Windows.Command = windowsCommand0
			tasksJSON.Tasks[index].Osx.Command = linuxCommand0
			tasksJSON.Tasks[index].Linux.Command = linuxCommand0

		}

		tasksJSONData, err := json.MarshalIndent(tasksJSON, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		workspaceTasksJSONFilePath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, "/.vscode/tasks.json")
		workspaceTasksJSONFile, err := os.OpenFile(workspaceTasksJSONFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer workspaceTasksJSONFile.Close()
		_, err = workspaceTasksJSONFile.Write(tasksJSONData)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		utils.CopyDir(goScriptsSourceDirPath, goScriptsDestDirPath)
	}

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


