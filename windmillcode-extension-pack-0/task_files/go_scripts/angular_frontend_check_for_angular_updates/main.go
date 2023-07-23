package main

import (
	"fmt"
	"go_scripts/utils"
	"os"
	"path/filepath"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToAngularApp()
	utils.RunCommand("npx",[]string{"ng","update"})
}

