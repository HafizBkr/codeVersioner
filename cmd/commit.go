package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func SaveCommit(message string) {
	repoDir := ".vcs"
	objectsDir := filepath.Join(repoDir, "objects")
	stagingDir := filepath.Join(repoDir, "staging")

	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		fmt.Println("Error: Repository not initialized. Run 'vcs start' first.")
		return
	}

	files, err := ioutil.ReadDir(stagingDir)
	if err != nil {
		fmt.Println("Error reading staging area:", err)
		return
	}

	commitID := fmt.Sprintf("%d", time.Now().UnixNano())
	commitDir := filepath.Join(objectsDir, commitID)
	os.MkdirAll(commitDir, 0755)

	for _, file := range files {
		src := filepath.Join(stagingDir, file.Name())
		dst := filepath.Join(commitDir, file.Name())

		content, err := ioutil.ReadFile(src)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		err = ioutil.WriteFile(dst, content, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			continue
		}
	}

	commitMetadata := fmt.Sprintf("Commit: %s\nMessage: %s\nDate: %s\n", commitID, message, time.Now().Format(time.RFC1123))
	ioutil.WriteFile(filepath.Join(objectsDir, commitID+".txt"), []byte(commitMetadata), 0644)

	fmt.Println("Commit saved:", commitID)

	os.RemoveAll(stagingDir)
	os.MkdirAll(stagingDir, 0755)
}
