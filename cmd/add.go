package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func AddFile(filename string) {
	repoDir := ".vcs"
	stagingDir := filepath.Join(repoDir, "staging")

	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		fmt.Println("Error: Repository not initialized. Run 'vcs start' first.")
		return
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(stagingDir, filepath.Base(filename)), content, 0644)
	if err != nil {
		fmt.Println("Error adding file:", err)
		return
	}

	fmt.Println("File added to staging area:", filename)
}
