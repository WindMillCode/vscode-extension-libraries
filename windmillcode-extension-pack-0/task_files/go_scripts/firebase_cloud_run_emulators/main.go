package main

import (
	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToFirebaseApp()

	utils.RunCommand("yarn", []string{"cleanup"})
	utils.RunCommand("npx", []string{"firebase", "emulators:start", "--import=devData", "--export-on-exit"})
}
