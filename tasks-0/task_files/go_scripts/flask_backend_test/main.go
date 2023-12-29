package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceFolder, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	settings, err := utils.GetSettingsJSON(workspaceFolder)
	if err != nil {
		return
	}
	utils.CDToFlaskApp()
	flaskAppFolder, err := os.Getwd()
	if err != nil {
		return
	}
	envVarsFile := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"where are the env vars located"},
			Default: utils.JoinAndConvertPathToOSFormat(workspaceFolder, settings.ExtensionPack.FlaskBackendTestHelperScript),
		},
	)
	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide a python version for pyenv to use"},
			Default: settings.ExtensionPack.PythonVersion0,
		},
	)
	if pythonVersion != "" {
		utils.RunCommand("pyenv", []string{"global", pythonVersion})
	}
	utils.CDToLocation(workspaceFolder)

	envVarCommandOptions := utils.CommandOptions{
		Command:      "windmillcode_go",
		Args:         []string{"run", envVarsFile, filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)), workspaceFolder},
		GetOutput:    true,
		TargetDir:     filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)),
	}
	envVars,err := utils.RunCommandWithOptions(envVarCommandOptions)
	if err != nil {
		return
	}
	envVarsArray := strings.Split(envVars, ",")


	for _, x := range envVarsArray {
		keyPair := []string{}
		for _, y := range strings.Split(x, "=") {
			keyPair = append(keyPair, strings.TrimSpace(y))
		}
		keyPair[1] = strings.ReplaceAll(keyPair[1], ",", "")
		os.Setenv(keyPair[0], keyPair[1])
	}
	flaskAppUnitTestFolder := utils.JoinAndConvertPathToOSFormat(
		flaskAppFolder, "unit_tests",
	)
	utils.CDToLocation(flaskAppUnitTestFolder)
	for {
		utils.RunCommand("python", []string{"run_tests.py"})
	}

}


