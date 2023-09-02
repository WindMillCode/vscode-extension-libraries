package main

import (
	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.CDToFlutterApp()

	utils.RunCommand("dart", []string{"fix","--apply"})
	utils.RunCommand("flutter", []string{"build","appbundle"})
}
