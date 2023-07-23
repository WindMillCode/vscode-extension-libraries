package utils

import (

	"io/ioutil"
	"os"
)

func OverwriteFile(filePath string, content []byte) error {
	return ioutil.WriteFile(filePath, content, 0644)
}

func FolderExists(path string) bool {
	_, err := os.Stat(path)
	if ; !os.IsNotExist(err) {
		// path/to/whatever exists
		return true
	} else{
		return false
	}
}
