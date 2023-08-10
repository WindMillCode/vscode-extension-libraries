package main

import (
	"strings"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRooot()
	utils.RunCommand("git", []string{"add", "."})
	utils.RunCommand("git", []string{"commit", "-m", "[CHECKPOINT] before upgrading to next angular version"})
	utils.CDToAngularApp()

	inputText := utils.RunCommandAndGetOutput("npx", []string{"ng", "update"})
	inputLines := strings.Split(inputText, "\n")
	packagesToUpdate := []string{}
	for _, line := range inputLines {
		if strings.Contains(line, "ng update @") {
			packagesToUpdate = append(packagesToUpdate, line)
		}
	}
	updateCommand := "ng update"
	for _, pkg := range packagesToUpdate {
		packageGroup := strings.TrimSpace(strings.Split(pkg, "->")[0])
		packageName := strings.TrimSpace(strings.Split(packageGroup, " ")[0])
		updateCommand += " " + packageName
	}
	utils.RunCommand("npx", strings.Split(updateCommand, " "))
	cliInfo := utils.ShowMenuModel{
		Prompt: "auto update additiona packages",
		Choices:[]string{"TRUE","FALSE"},
	}
	addtl := utils.ShowMenu(cliInfo,nil)
	if addtl == "TRUE"{
		cliInfo := utils.ShowMenuModel{
			Prompt: "choose the package manager",
			Choices:[]string{"npm","yarn"},
			Default:"npm",
		}
		packageManager := utils.ShowMenu(cliInfo,nil)
		if packageManager == "yarn"{
			utils.RunCommand("yarn", []string{"upgrade", "@windmillcode/angular-wml-components-base", "@rxweb/reactive-form-validators", "@fortawesome/fontawesome-free", "@compodoc/compodoc", "@sentry/angular-ivy", "@sentry/tracing"})
			utils.RunCommand("yarn", []string{"upgrade", "--dev", "@faker-js/faker", "@windmillcode/angular-templates", "webpack-bundle-analyzer", "browserify"})
		} else {
			utils.RunCommand("npm", []string{"update", "@windmillcode/angular-wml-components-base", "@rxweb/reactive-form-validators", "@fortawesome/fontawesome-free", "@compodoc/compodoc", "@sentry/angular-ivy", "@sentry/tracing"})
			utils.RunCommand("npm", []string{"update", "--dev", "@faker-js/faker", "@windmillcode/angular-templates", "webpack-bundle-analyzer", "browserify"})
		}
	}


}
