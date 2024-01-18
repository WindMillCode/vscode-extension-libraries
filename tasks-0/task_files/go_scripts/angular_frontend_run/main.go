package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	angularFrontend := settings.ExtensionPack.AngularFrontend

	cliInfo := utils.ShowMenuModel{
		Prompt:  "run with cache?",
		Choices: []string{"true", "false"},
	}
	runWithCache := utils.ShowMenu(cliInfo, nil)

	defaultConfig := "development"
	if utils.IsRunningInDocker() {
		defaultConfig = strings.Replace(defaultConfig, "development", "docker-development", 1)
		for key, val := range angularFrontend.Configurations {
			angularFrontend.Configurations[key] = strings.Replace(val, "development", "docker-development", 1)
		}
	}

	cliInfo = utils.ShowMenuModel{
		Prompt:  "the configuration to run",
		Choices: angularFrontend.Configurations,
		Default: defaultConfig,
		Other:   true,
	}
	serveConfiguration := utils.ShowMenu(cliInfo, nil)

	utils.CDToAngularApp()
	if runWithCache == "false" {
		if err := os.RemoveAll(utils.JoinAndConvertPathToOSFormat(".", ".angular")); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	utils.RunCommand("npx", []string{"ng", "serve", "-c", serveConfiguration, "--ssl=true", fmt.Sprintf("--ssl-key=%s", os.Getenv("WML_CERT_KEY0")), fmt.Sprintf("--ssl-cert=%s", os.Getenv("WML_CERT0"))})
}
