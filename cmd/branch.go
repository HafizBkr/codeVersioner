package cmd

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func CreateBranch(name string) {
    repoDir := ".vcs"
    refsDir := filepath.Join(repoDir, "refs")

    if _, err := os.Stat(repoDir); os.IsNotExist(err) {
        fmt.Println("Error: Repository not initialized. Run 'vcs start' first.")
        return
    }

    // Vérifier si le dossier refs existe
    if _, err := os.Stat(refsDir); os.IsNotExist(err) {
        os.MkdirAll(refsDir, 0755)
    }

    // Créer le fichier de branche
    branchPath := filepath.Join(refsDir, name)
    
    // Lire le HEAD actuel
    headContent, err := ioutil.ReadFile(filepath.Join(repoDir, "HEAD"))
    if err != nil {
        fmt.Println("Error: Cannot read HEAD")
        return
    }

    // Écrire le HEAD dans le nouveau fichier de branche
    err = ioutil.WriteFile(branchPath, headContent, 0644)
    if err != nil {
        fmt.Println("Error creating branch:", err)
        return
    }

    fmt.Printf("Created branch '%s'\n", name)
}