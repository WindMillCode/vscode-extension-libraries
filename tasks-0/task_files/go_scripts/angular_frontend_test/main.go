package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	utils.CDToWorkspaceRoot()
	utils.CDToAngularApp()



	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		utils.RunCommand("npm", []string{"run", "test"})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		port := 8003
		utils.CDToLocation(utils.ConvertPathToOSFormat("coverage"))
		folderPath,_ := os.Getwd()
		files, err := os.ReadDir(folderPath)
		utils.CDToLocation(utils.ConvertPathToOSFormat(files[0].Name()))
		http.Handle("/", http.FileServer(http.Dir(".")))
		err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			fmt.Println("Error starting the server:", err)
		}

	}()
	wg.Wait()
}
