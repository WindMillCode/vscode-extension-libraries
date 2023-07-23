package main

import (
	"fmt"
	"go_scripts/utils"
	"os"
	"path/filepath"
)

func main() {

	utils.CDToWorkspaceRooot()
	cliInfo := utils.ShowMenuModel{
		Prompt: "Choose an option:",
		Choices:[]string{"dev", "preview", "prod"},
	}
	envType := utils.ShowMenu( cliInfo, nil)
	utils.CDToAngularApp()
	utils.RunCommand("yarn",[]string{"analyze:" + envType})
}

