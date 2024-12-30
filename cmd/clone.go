package cmd

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
)

func CloneLocalRepo(localPath string) {
    // Vérifier si le dépôt local existe
    if _, err := os.Stat(localPath); os.IsNotExist(err) {
        fmt.Println("Error: Local repository does not exist.")
        return
    }

    // Extraire le nom du dépôt à partir du chemin (en utilisant le dernier élément du chemin)
    repoName := filepath.Base(localPath)

    // Créer un répertoire pour le clone
    cloneDir := filepath.Join(".", repoName)
    if _, err := os.Stat(cloneDir); !os.IsNotExist(err) {
        fmt.Println("Error: Directory already exists:", cloneDir)
        return
    }

    // Utiliser git pour cloner le dépôt local dans un nouveau répertoire
    cmd := exec.Command("git", "clone", localPath, cloneDir)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        fmt.Println("Error cloning repository:", err)
        return
    }

    fmt.Printf("Repository '%s' cloned successfully into '%s'.\n", repoName, cloneDir)

    // Optionnel : Si vous souhaitez récupérer uniquement les fichiers poussés, vous pouvez effectuer un `git pull`
    // dans le répertoire cloné.
    pullCmd := exec.Command("git", "-C", cloneDir, "pull")
    pullCmd.Stdout = os.Stdout
    pullCmd.Stderr = os.Stderr
    err = pullCmd.Run()
    if err != nil {
        fmt.Println("Error pulling files:", err)
        return
    }

    fmt.Println("Successfully pulled the latest files.")
}
