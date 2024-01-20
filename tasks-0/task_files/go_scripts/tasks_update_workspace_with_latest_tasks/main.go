package main

import (
	"encoding/json"
	"fmt"
	"main/shared"
	"os"
	"regexp"
	"strings"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {
	workSpaceFolder := os.Args[1]
	extensionFolder := os.Args[2]
	tasksJsonRelativeFilePath := os.Args[3]
	settings, err := utils.GetSettingsJSON(workSpaceFolder)
	if err != nil {
		return
	}
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
	goScriptsSourceDirPath := utils.JoinAndConvertPathToOSFormat(extensionFolder, "task_files/go_scripts")
	goScriptsDestDirPath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, "ignore/Windmillcode/go_scripts")
	cliInfo = utils.ShowMenuModel{
		Prompt:  "delete dest dir to ensure proper update (if updates are not taking place choose YES)",
		Choices: []string{"YES", "NO"},
		Default: "YES",
	}
	deleteDestDir := utils.ShowMenu(cliInfo, nil)
	if deleteDestDir == "YES" {
		fmt.Println("Deleting Dest dir ...")
		if err := os.RemoveAll(goScriptsDestDirPath); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	fmt.Println("Copying over files ...")
	utils.CopyDir(goScriptsSourceDirPath, goScriptsDestDirPath)

	if proceed == "TRUE" {

		cliInfo := utils.ShowMenuModel{
			Prompt:  "Run Tasks Via Interpreted or Complied Mode",
			Choices: []string{"COMPLIED", "INTERPRETED"},
			Default: "COMPLIED",
		}
		runMode := utils.ShowMenu(cliInfo, nil)

		cliInfo = utils.ShowMenuModel{
			Prompt:  "use default user (if unsure select NO)",
			Choices: []string{"NO", "YES"},
			Default: "NO",
		}
		customUserIsPresent := utils.ShowMenu(cliInfo, nil)

		for index, task := range tasksJSON.Tasks {

			pattern0 := ":"
			regex0 := regexp.MustCompile(pattern0)
			programLocation0 := regex0.Split(task.Label, -1)
			pattern1 := " "
			regex1 := regexp.MustCompile(pattern1)
			programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
			programLocation2 := strings.Join(programLocation1, "_")
			programLocation3 := "ignore//${input:current_user_0}//go_scripts//" + programLocation2
			if customUserIsPresent == "NO" {
				programLocation3 = "ignore//Windmillcode//go_scripts//" + programLocation2
			}
			taskExecutable := ".//main"
			if runMode == "INTERPRETED" {
				taskExecutable = fmt.Sprintf("%s %s", goExecutable, "run main.go")
			}
			linuxCommand0 := "cd " + programLocation3 + " ; " + taskExecutable
			windowsCommand0 := "cd " + strings.Replace(programLocation3, "//", "\\", -1) + " ; " + strings.Replace(taskExecutable, "//", "\\", -1)
			tasksJSON.Tasks[index].Windows.Command = windowsCommand0
			tasksJSON.Tasks[index].Osx.Command = linuxCommand0
			tasksJSON.Tasks[index].Linux.Command = linuxCommand0
			tasksJSON.Tasks[index].Linux.Options = shared.CommandOptions{
				Shell: shared.ShellOptions{
					Executable: "bash",
					Args:       []string{"-ic"},
				},
			}
			if strings.Contains(strings.Join(settings.ExtensionPack.TasksToRunOnFolderOpen, " , "), task.Label) {
				tasksJSON.Tasks[index].RunOptions.RunOn = "folderOpen"
			}

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
	}

	shared.RebuildExecutables(proceed, cliInfo, tasksJSON, goScriptsDestDirPath, goExecutable)

}
