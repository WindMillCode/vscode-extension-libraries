package main

import (
	"go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToAngularApp()
	utils.RunCommand("npx",[]string{"ng","update"})
}
