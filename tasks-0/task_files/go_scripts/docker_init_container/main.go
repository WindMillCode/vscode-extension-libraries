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

	initOptions := utils.CommandOptions{
		Command: "windmillcode_go",
		Args: []string{"run", initScript, initScriptArgs},
	}
	_,err :=utils.RunCommandWithOptions(initOptions)
	if err != nil {
		initOptions = utils.CommandOptions{
			Command: "go",
			Args: []string{"run", initScript, initScriptArgs},
		}
		_,err =utils.RunCommandWithOptions(initOptions)
		if err != nil {
			return
		}
	}
}
