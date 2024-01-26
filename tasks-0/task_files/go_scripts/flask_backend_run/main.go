package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/windmillcode/go_cli_scripts/v4/utils"
)

func main() {

	// scriptDir,_ := os.Getwd()
	// helperScript := utils.JoinAndConvertPathToOSFormat(scriptDir,"run_script.go")
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

	helperScript := settings.ExtensionPack.FlaskBackendDevHelperScript
	if utils.IsRunningInDocker() {
		helperScript = strings.Replace(helperScript, "dev", "docker_dev", 1)
	}

	cliInfo := utils.ShowMenuModel{
		Prompt: "Where are the env vars located",
		Choices: []string{
			utils.JoinAndConvertPathToOSFormat(workspaceFolder, helperScript),
		},
		Other: true,
	}
	envVarsFile := utils.ShowMenu(cliInfo, nil)

	pythonVersion := utils.GetInputFromStdin(
		utils.GetInputFromStdinStruct{
			Prompt:  []string{"provide a python version for pyenv to use"},
			Default: settings.ExtensionPack.PythonVersion0,
		},
	)

	if pythonVersion != "" {
		utils.RunCommand("pyenv", []string{"global", pythonVersion})
	}
	for {
		utils.CDToLocation(workspaceFolder)
		envVarCommandOptions := utils.CommandOptions{
			Command:     "windmillcode_go",
			Args:        []string{"run", envVarsFile, filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)), workspaceFolder},
			GetOutput:   true,
			TargetDir:   filepath.Dir(utils.JoinAndConvertPathToOSFormat(envVarsFile)),
			PrintOutput: false,
		}
		envVars, err := utils.RunCommandWithOptions(envVarCommandOptions)
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

		flaskAppCommand := utils.CommandOptions{
			Command:            "python",
			Args:               []string{"app.py", "--use_reloader", "False"},
			GetOutput:          false,
			PrintOutputOnly:    true,
			NonBlocking:        true,
			TargetDir:          flaskAppFolder,
			IsInputFromProgram: true,
		}
		flaskAppCommand.Self = &flaskAppCommand

		flaskDirOptions := utils.WatchDirectoryParams{
			DirectoryToWatch: flaskAppFolder,
			DebounceInMs:     5000,
			StartOnWatch:     true,
			ExcludePatterns:  []string{"**\\site-packages\\**"},
			Predicate: func(event fsnotify.Event) {
				content, err := os.ReadFile(event.Name)
				if err != nil {
						log.Fatal(err)
				}

				fmt.Println(string(content))


				self := flaskAppCommand.Self
				if self != nil && self.CmdObj != nil {
					self.CmdObj.Process.Kill()
					// pid := self.CmdObj.Process.Pid
					// if runtime.GOOS == "windows" {
					// 	utils.RunCommandWithOptions(utils.CommandOptions{
					// 		Command: "taskkill",
					// 		Args:    []string{"/F", "/PID", fmt.Sprint(pid)},
					// 	})
					// } else {
					// 	utils.RunCommandWithOptions(utils.CommandOptions{
					// 		Command: "pkill",
					// 		Args:    []string{"-9", fmt.Sprint(pid)},
					// 	})
					// }
				}

				runFlaskApp(flaskAppCommand)
			},
		}
		WatchDirectory(
			flaskDirOptions,
		)

	}

}

func runFlaskApp(flaskAppCmd utils.CommandOptions) {
	utils.RunCommandWithOptions(flaskAppCmd)
}

func WatchDirectory(options utils.WatchDirectoryParams) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Compile glob patterns

	// Function to check if a path should be included or excluded
	shouldIncludePath := func(path string) bool {
		return !(strings.Contains(path, "site-packages") || strings.Contains(path, "__pycache__") || strings.Contains(path, "unit_tests") || strings.HasSuffix(path, ".pyc"))

	}

	// Call the Predicate immediately for each file if StartOnWatch is true
	if options.StartOnWatch {
		// Manually creating a test event
		testEvent := fsnotify.Event{Name: options.DirectoryToWatch, Op: fsnotify.Write}
		options.Predicate(testEvent)
	}

	// Setup the watcher
	if err := filepath.Walk(options.DirectoryToWatch,
		func(path string, fi os.FileInfo, err error) error {
			if shouldIncludePath(path) {
				if fi.Mode().IsRegular() {
					return watcher.Add(path)
				}
			}

			return nil
		},
	); err != nil {
		fmt.Println("ERROR", err)
	}

	fmt.Printf("Watching directory %s\n", options.DirectoryToWatch)

	done := make(chan bool)

	go func() {
		var lastEventTime time.Time

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// Calculate the time elapsed since the last event
				elapsedTime := time.Since(lastEventTime)
				if int(elapsedTime.Milliseconds()) >= options.DebounceInMs {
					options.Predicate(event)
				}

				// Update the last event time
				lastEventTime = time.Now()

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}
