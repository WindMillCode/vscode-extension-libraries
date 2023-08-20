package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	// "sync"

	"github.com/windmillcode/go_scripts/utils"
)

func main() {


	dockerImages := utils.RunCommandAndGetOutput("docker", []string{"images", "--format", "'{{.Repository}}:{{.Tag}}'", "-a"})
	dockerImagesArray := strings.Split(dockerImages, "\n")

	var wg sync.WaitGroup
	for _, repoTag := range dockerImagesArray {
		if !strings.Contains(repoTag,"<none>") {
			wg.Add(1)
			fmt.Println(repoTag)
			repoTagWithoutQuotes := strings.Replace(repoTag,"'","",-1)
			regex0 := regexp.MustCompile(`[^a-zA-Z0-9|\.]`)
			tarImageName := regex0.ReplaceAllString(repoTagWithoutQuotes,"_")
			fmt.Println(tarImageName)
			go func(){
				defer wg.Done()
				utils.RunCommand("docker",[]string{"save","-o",fmt.Sprintf("%s.tar",tarImageName),repoTagWithoutQuotes});
			}()
		}
	}
	wg.Wait()
}

// docker images --format '{{.Repository}}:{{.Tag}}' -a
