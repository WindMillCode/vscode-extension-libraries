package utils

import (
	"os"
	"path/filepath"

)

func CDToLocation(location string){
	if err := os.Chdir(location); err != nil {
		panic(err)
	}
}
func CDToWorkspaceRooot(){
	CDToLocation(filepath.Join("..","..","..",".."))
}
func CDToAngularApp(){
	CDToLocation(filepath.Join("apps","frontend","AngularApp"))
}

func CDToFirebaseApp(){
	CDToLocation(filepath.Join("apps","cloud","FirebaseApp"))
}
