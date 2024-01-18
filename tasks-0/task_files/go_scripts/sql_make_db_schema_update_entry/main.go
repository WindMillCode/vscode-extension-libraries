package main

import (
	"fmt"
	"os"
	"time"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		return
	}
	settings, err := utils.GetSettingsJSON(workspaceRoot)
	if err != nil {
		return
	}
	cliInfo := utils.ShowMenuModel{
		Prompt:  "Enter the database script location (refer to the folder in apps\\database for your project)",
		Choices: settings.ExtensionPack.DatabaseOptions,
		Default: "mysql",
	}
	databaseToBackup := utils.ShowMenu(cliInfo, nil)
	databaseBackupLocation := utils.JoinAndConvertPathToOSFormat("apps", "database", databaseToBackup, "schema_entries")

	myEnvs := []string{"dev", "preview", "prod"}
	for _, v := range myEnvs {
		utils.CDToLocation(workspaceRoot)
		utils.CDToLocation(databaseBackupLocation)
		utils.CDToLocation(utils.JoinAndConvertPathToOSFormat(v))
		currentDay := time.Now().Format("06-1-02_03-04-05")
		err := os.MkdirAll(currentDay, 0755)
		if err != nil {
			fmt.Printf("unable to make %s in %s: \n Err msg: %s", currentDay, v, err.Error())
		}
		utils.CopySelectFilesToDestination(
			utils.CopySelectFilesToDestinationStruct{
				GlobPattern:    "*.sql",
				DestinationDir: currentDay,
			},
		)

	}

}
