package shared

import (
	"fmt"
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

func RebuildExecutables(proceed string, cliInfo utils.ShowMenuModel, tasksJSON TasksJSON, goScriptsDestDirPath string, goExecutable string) {
	var rebuild string
	if proceed == "TRUE" {
		rebuild = "TRUE"
	} else {
		cliInfo = utils.ShowMenuModel{
			Prompt:  "Do you want to rebuild the go programs into binary exectuables ",
			Choices: []string{"TRUE", "FALSE"},
		}
		rebuild = utils.ShowMenu(cliInfo, nil)
	}

	if rebuild == "TRUE" {

		cliInfo = utils.ShowMenuModel{
			Prompt:  "Build one by one or all at once",
			Choices: []string{"ALL_AT_ONCE", "ONE_BY_ONE"},
		}
		inSync := utils.ShowMenu(cliInfo, nil)

		if inSync == "ALL_AT_ONCE" {
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
					BuildGoCLIProgram(absProgramLocation, goExecutable)
				}()
			}
			wg.Wait()
		} else{
			fmt.Print(len(tasksJSON.Tasks))
			for _, task := range tasksJSON.Tasks {

				pattern0 := ":"
				regex0 := regexp.MustCompile(pattern0)
				programLocation0 := regex0.Split(task.Label, -1)
				pattern1 := " "
				regex1 := regexp.MustCompile(pattern1)
				programLocation1 := regex1.Split(strings.Join(programLocation0, ""), -1)
				programLocation2 := strings.Join(programLocation1, "_")
				absProgramLocation := utils.JoinAndConvertPathToOSFormat(goScriptsDestDirPath, programLocation2)
				BuildGoCLIProgram(absProgramLocation, goExecutable)

			}
		}

	}
}

func BuildGoCLIProgram(programLocation string, goExecutable string) {

	fmt.Printf("%s \n", programLocation)
	utils.RunCommandInSpecificDirectory(goExecutable, []string{"build", "main.go"}, programLocation)
	fmt.Printf("Finished building %s \n", programLocation)

}
