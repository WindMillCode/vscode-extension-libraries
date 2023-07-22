package main

import (
	"fmt"
	"os"
	// "windmillcode_tasks_json_scripts/bubbletea2"
	"../utils"
)

func main() {
	options := []string{"dev", "preview", "prod"}
	cliInfo := utils.ShowMenuModel{
		Prompt: "Choose an option:",
		Choices:options,
	}
	envType := utils.ShowMenu( cliInfo, nil)
	fmt.Printf(envType)
	cwd,_ := os.Getwd()
	fmt.Printf(cwd)
}
