package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

type Task struct {
	Label   string `json:"label"`
	Type    string `json:"type"`
	Windows struct {
		Command string `json:"command"`
	} `json:"windows"`
	Linux struct {
		Command string `json:"command"`
	} `json:"linux"`
	Osx struct {
		Command string `json:"command"`
	} `json:"osx"`
	RunOptions struct {
		RunOn         string `json:"runOn"`
		InstanceLimit int    `json:"instanceLimit"`
	} `json:"runOptions"`
}

type Input struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Default     string `json:"default"`
	Type        string `json:"type"`
}
type TasksJSON struct {
	Version string  `json:"version"`
	Tasks   []Task  `json:"tasks"`
	Inputs  []Input `json:"inputs"`
}

func main() {
	utils.CDToWorkspaceRoot()
	workSpaceFolder,_ :=os.Getwd()
	cliInfo := utils.ShowMenuModel{
		Prompt: "choose the executable to use (try with windmillcode_go first if not then use go)",
		Choices:[]string{"go","windmillcode_go"},
		Default:  "windmillcode_go",
	}
	goExecutable := utils.ShowMenu(cliInfo,nil)

	proceed := "FALSE"

	tasksJsonFilePath := utils.JoinAndConvertPathToOSFormat(workSpaceFolder, ".vscode/tasks.json");

	content, err, shouldReturn := createTasksJson(tasksJsonFilePath, false)
	if shouldReturn {
		return
	}

	var tasksJSON TasksJSON
	err = json.Unmarshal(content, &tasksJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	goScriptsDestDirPath := filepath.Join(workSpaceFolder, "ignore/Windmillcode/go_scripts")

	rebuildExecutables(proceed, cliInfo, tasksJSON, goScriptsDestDirPath, goExecutable)

}

func rebuildExecutables(proceed string, cliInfo utils.ShowMenuModel, tasksJSON TasksJSON, goScriptsDestDirPath string, goExecutable string) {
	var rebuild string
	rebuild = "TRUE"

	if rebuild == "TRUE" {
		var wg sync.WaitGroup
		fmt.Print(len(tasksJSON.Tasks))
		for _, task := range tasksJSON.Tasks {
			wg.Add(1)

			pattern0 := ":"
			regex0 := regexp.MustCompile(pattern0)
			programLocation0 := regex0.Split(task.Label, -1)
			pattern1 := " "
			regex1 := regexp.MustCompile(pattern1)
			programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
			programLocation2 := strings.Join(programLocation1, "_")
			absProgramLocation := utils.JoinAndConvertPathToOSFormat(goScriptsDestDirPath, programLocation2)
			go func() {
				defer wg.Done()
				buildGoCLIProgram(absProgramLocation, goExecutable)
			}()
		}
		wg.Wait()
	}
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

func buildGoCLIProgram(programLocation string, goExecutable string) {

	fmt.Printf("%s \n", programLocation)
	utils.RunCommandInSpecificDirectory(goExecutable, []string{"build", "main.go"}, programLocation)
	fmt.Printf("Finished building %s \n", programLocation)

}
