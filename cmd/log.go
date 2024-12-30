package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ShowCommits() {
	repoDir := ".vcs"
	objectsDir := filepath.Join(repoDir, "objects")

	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		fmt.Println("Error: Repository not initialized. Run 'vcs start' first.")
		return
	}

	files, err := ioutil.ReadDir(objectsDir)
	if err != nil {
		fmt.Println("Error reading objects directory:", err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			content, err := ioutil.ReadFile(filepath.Join(objectsDir, file.Name()))
			if err != nil {
				fmt.Println("Error reading commit file:", err)
				continue
			}
			fmt.Println(string(content))
		}
	}
}
