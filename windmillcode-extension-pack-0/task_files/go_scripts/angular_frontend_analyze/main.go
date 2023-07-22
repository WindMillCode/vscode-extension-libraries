package main

import (
	"fmt"
	"go_scripts/utils"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	if err := os.Chdir(filepath.Join("..","..","..","..")); err != nil {
		fmt.Println("Error:", err)
		return
	}
	cwd, err := os.Getwd()
	if  err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current Working Directory:", cwd)
	options := []string{"dev", "preview", "prod"}
	cliInfo := utils.ShowMenuModel{
		Prompt: "Choose an option:",
		Choices:options,
	}
	envType := utils.ShowMenu( cliInfo, nil)
	if err := os.Chdir(filepath.Join("apps","frontend","AngularApp")); err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.RunCommand("yarn",[]string{"analyze:" + envType})
}

