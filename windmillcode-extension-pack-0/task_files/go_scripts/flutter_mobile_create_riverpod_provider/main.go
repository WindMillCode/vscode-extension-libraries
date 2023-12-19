package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/windmillcode/go_cli_scripts/v3/utils"
)

func main() {

	scriptLocation, err := os.Getwd()
	if err != nil {
		return
	}
	templateLocation := utils.JoinAndConvertPathToOSFormat(scriptLocation, "template")
	utils.CDToWorkspaceRoot()
	utils.CDToFlutterApp()
	flutterApp, err := os.Getwd()
	if err != nil {
		return
	}
	providerName := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt: []string{"The name you would like to give to the provider"},
			ErrMsg: "You must provide a value",
		},
	)

	snakeCaseProviderName := strcase.ToSnake(providerName)
	providerLocation := utils.JoinAndConvertPathToOSFormat(flutterApp, "lib", "util", "riverpod_providers", snakeCaseProviderName)
	newTemplatePath := utils.JoinAndConvertPathToOSFormat(providerLocation, fmt.Sprintf("%s.dart", snakeCaseProviderName))
	utils.CopyDir(templateLocation, providerLocation)
	os.Rename(
		utils.JoinAndConvertPathToOSFormat(providerLocation, "template.dart"),
		newTemplatePath,
	)

	fileString, err := utils.ReadFile(newTemplatePath)
	if err != nil {
		return
	}
	fileString = strings.ReplaceAll(fileString, "Template", strcase.ToCamel(providerName))
	fileString = strings.ReplaceAll(fileString, "Wml", "WML")
	fileString = strings.ReplaceAll(fileString, "template", strings.ToLower(providerName))
	utils.OverwriteFile(newTemplatePath, fileString)
}
