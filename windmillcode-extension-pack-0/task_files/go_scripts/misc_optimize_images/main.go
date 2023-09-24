package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	workspaceRoot, err := os.Getwd()
	if err != nil {
		fmt.Println("there was an error while trying to receive the current dir")
	}
	projectsCLIString := utils.TakeVariableArgs(
		utils.TakeVariableArgsStruct{
			Prompt:  "Provide the paths of all the locations where you want your images optimized",
			Default: workspaceRoot,
		},
	)
	backupLocation := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"This program will delete images in the directories provided please provide a path "},
			Default: "",
		},
	)
	optimizePercent := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"enter a value from 1 -100 where 100 is perform no changes and 0 is full optimization, recommnded is"},
			Default: "75",
		},
	)

	var wg sync.WaitGroup
	regex0 := regexp.MustCompile(" ")
	projectsList := regex0.Split(projectsCLIString, -1)
	for _, project := range projectsList {
		app := filepath.Join(project)
		normalizedBackupLocation := ""
		// if runtime.GOOS == "windows"{
		// 	normalizedBackupLocation = filepath.Join(backupLocation,utils.RemoveDrivePath(app))
		// } else {
		// }
		normalizedBackupLocation = filepath.Join(backupLocation, app)
		wg.Add(1)
		go func() {
			defer wg.Done()
			utils.CopyDir(app, normalizedBackupLocation)
			allEntries, err := utils.GetItemsInFolderRecursive(app, true)
			if err != nil {
				fmt.Println("An error occured while recursively goin through the directory", err)
			}
			for _, entry := range allEntries {
				prefixImage := utils.HasSuffixInArray(entry, []string{".png", ".gif", ".ico", ".jpg", ".svg", ".webp", ".ico"}, true)
				if prefixImage != "" {
					imageFolderPath := filepath.Dir(entry)
					imageFile := filepath.Base(entry)
					wg.Add(1)
					go func() {
						defer wg.Done()
						utils.RunCommandInSpecificDirectory("convert", []string{"-quality", optimizePercent, imageFile, fmt.Sprintf("%s%s", prefixImage, ".jpg")}, imageFolderPath);
						os.Remove(entry)
					}()
				}
			}

		}()
	}
	wg.Wait()

}
