package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Status() {
    repoDir := ".vcs"
    stagingDir := filepath.Join(repoDir, "staging")

    if _, err := os.Stat(repoDir); os.IsNotExist(err) {
        fmt.Println("Error: Repository not initialized. Run 'vcs start' first.")
        return
    }

    fmt.Println("On branch master")

    // Vérifier les fichiers stagés
    files, err := ioutil.ReadDir(stagingDir)
    if err != nil {
        fmt.Println("Error reading staging area:", err)
        return
    }

    if len(files) == 0 {
        fmt.Println("\nNo changes staged for commit")
        return
    }

    fmt.Println("\nChanges staged for commit:")
    for _, file := range files {
        fmt.Printf("\tmodified: %s\n", file.Name())
    }
}