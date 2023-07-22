package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

func OverwriteFile(filePath string, content []byte) error {
	return ioutil.WriteFile(filePath, content, 0644)
}

// getType returns the type of a given value as a string
func GetType(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func TakeVariableArgs() []string {
	var innerScriptArguments []string

	fmt.Println("Enter the arguments to pass to the script (press ENTER to enter another argument, leave blank and press ENTER once done):")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		argument := strings.TrimSpace(scanner.Text())
		if argument == "" {
			break
		}
		innerScriptArguments = append(innerScriptArguments, argument)
	}

	return innerScriptArguments
}

func GetParamValue(parameterName string, parameterValue interface{}) interface{} {
	if parameterValue != nil {
		return parameterValue
	} else {
		fmt.Printf("Parameter '%s' value not found.\n", parameterName)
		return nil
	}
}

func GetCurrentPath() string {
	executablePath, err := os.Executable()
	if err != nil {
		// Handle the error if necessary
		return ""
	}
	return filepath.Dir(executablePath)
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Remove leading/trailing whitespaces and newline characters
	branch := strings.TrimSpace(string(output))

	return branch, nil
}


// Function to clear the console screen
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
