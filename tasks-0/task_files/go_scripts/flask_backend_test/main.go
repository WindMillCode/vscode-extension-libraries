package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
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
		TargetDir:    filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)),
		PanicOnError: false,
	}
	envVars, err := utils.RunCommandWithOptions(envVarCommandOptions)
	if err != nil {
		envVarCommandOptions := utils.CommandOptions{
			Command:   "go",
			Args:      []string{"run", envVarsFile, filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)), workspaceFolder},
			GetOutput: true,
			TargetDir: filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)),
		}
		envVars, err = utils.RunCommandWithOptions(envVarCommandOptions)
		if err != nil {
			return
		}
	}
	envVarsArray := strings.Split(envVars, ",")

	for _, x := range envVarsArray {
		keyPair := []string{}
		for _, y := range strings.Split(x, "=") {
			keyPair = append(keyPair, strings.TrimSpace(y))
		}
		os.Setenv(keyPair[0], keyPair[1])
	}
	flaskAppUnitTestFolder := utils.JoinAndConvertPathToOSFormat(
		flaskAppFolder, "unit_tests",
	)
	utils.CDToLocation(flaskAppUnitTestFolder)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		port := 8004
		utils.CDToLocation(utils.JoinAndConvertPathToOSFormat("covhtml"))
		http.Handle("/", http.FileServer(http.Dir(".")))
		err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			fmt.Println("Error starting the server:", err)
		}
	}()


	wg.Add(1)
	go func() {
		defer wg.Done()
		currentDir,_ := os.Getwd()
		testsDir := utils.JoinAndConvertPathToOSFormat(currentDir)
		runTestCases(testsDir)
		// watchDirectory(testsDir,10,[]string{"covhtml","__pycache__",".pytest_cache",".coverage"})

	}()


	wg.Wait()

}


func watchDirectory(directoryToWatch string, backoff int, ignoredDirs []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	if err := filepath.Walk(directoryToWatch,
		func(path string, fi os.FileInfo, err error) error {
			if fi.Mode().IsDir() && !shouldIgnoreDir(path, ignoredDirs) {
				return watcher.Add(path)
			}
			return nil
		},
	); err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Printf("Watching directory %s\n", directoryToWatch)

	done := make(chan bool)

	go func() {
		var lastEventTime time.Time

		for {
			select {
				case _ = <-watcher.Events:

					// Calculate the time elapsed since the last event
					elapsedTime := time.Since(lastEventTime)
					if int(elapsedTime.Seconds()) >= backoff {
						runTestCases(directoryToWatch)
						fmt.Println("finished")
					}

					// Update the last event time
					lastEventTime = time.Now()

				case err := <-watcher.Errors:
					fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func shouldIgnoreDir(dirPath string, ignoredDirs []string) bool {
	// Check if the dirPath matches any of the ignored directories
	for _, ignoredDir := range ignoredDirs {
		if strings.HasSuffix(dirPath, ignoredDir) {
			return true
		}
	}
	return false
}
func runTestCases(directoryToWatch string) {
	testOptions := utils.CommandOptions{
		Command:   "python",
		Args:      []string{"run_tests.py"},
		TargetDir: directoryToWatch,
	}

	utils.RunCommandWithOptions(testOptions)
}
