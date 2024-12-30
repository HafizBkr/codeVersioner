package cmd

import (
	"fmt"
	"os"
)

func InitRepo() {
	repoDir := ".vcs"
	subDirs := []string{"staging", "objects"}

	if _, err := os.Stat(repoDir); !os.IsNotExist(err) {
		fmt.Println("Repository already initialized.")
		return
	}

	if err := os.Mkdir(repoDir, 0755); err != nil {
		fmt.Println("Error initializing repository:", err)
		return
	}

	for _, subDir := range subDirs {
		os.Mkdir(fmt.Sprintf("%s/%s", repoDir, subDir), 0755)
	}

	fmt.Println("Repository initialized.")
}
