package main

import (
	"fmt"
	"path/filepath"

	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	utils.CDToWorkspaceRoot()

	initScript := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"docker init script to run relative to workspace root "},
			Default: utils.JoinAndConvertPathToOSFormat("ignore\\Local\\docker_init_container.go"),
		},
	)
	initScriptArgs := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{},
	)
	initScriptArgs = fmt.Sprintf("%s %s", utils.JoinAndConvertPathToOSFormat("..", "..", ".."), initScriptArgs)
	initScriptLocation := filepath.Dir(initScript)
	utils.CDToLocation(initScriptLocation)
	initScript = filepath.Base(initScript)

	utils.RunCommand("windmillcode_go", []string{"run", initScript, initScriptArgs})
}
