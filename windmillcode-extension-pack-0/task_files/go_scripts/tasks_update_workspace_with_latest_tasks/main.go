package main

import (
	"fmt"
	"os"
	"path"
	"go_scripts/utils"
)

func main() {
	workSpaceFolder := os.Args[1]
	extensionFolder := os.Args[2]
	tasksJsonFilePath := os.Args[3]
	cliInfo := utils.ShowMenuModel{
		Prompt: "Replace tasks.json?",
		Choices:[]string{"TRUE", "FALSE"},
	}
	proceed := utils.ShowMenu( cliInfo, nil)
	if proceed == "FALSE" {
		return
	}
	joinedPath := path.Join(extensionFolder, tasksJsonFilePath)

	fmt.Println("Joined path:", joinedPath)
	file, err := os.Open(joinedPath)
	if err != nil {
		fmt.Println("Error opening the JSON file:", err)
		return
	}
	defer file.Close()

	err = os.Chdir(path.Join(workSpaceFolder, "/.vscode"))
	if err != nil {
		fmt.Println("Error changing the working directory:", err)
		return
	}



}
