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

    // Vérifier si le répertoire .vcs existe
    if _, err := os.Stat(repoDir); os.IsNotExist(err) {
        fmt.Println("Error: Repository not initialized. Run 'vcs start' first.")
        return
    }

    // Vérifier si le répertoire refs existe, sinon, le créer
    if _, err := os.Stat(refsDir); os.IsNotExist(err) {
        os.MkdirAll(refsDir, 0755)
    }

    // Créer le chemin de la nouvelle branche
    branchPath := filepath.Join(refsDir, name)

    // S'assurer que le répertoire pour la branche existe
    branchDir := filepath.Dir(branchPath)
    if _, err := os.Stat(branchDir); os.IsNotExist(err) {
        os.MkdirAll(branchDir, 0755)
    }

    // Lire le contenu du HEAD actuel
    headContent, err := ioutil.ReadFile(filepath.Join(repoDir, "HEAD"))
    if err != nil {
        fmt.Println("Error: Cannot read HEAD")
        return
    }

    // Écrire le HEAD dans le fichier de la nouvelle branche
    err = ioutil.WriteFile(branchPath, headContent, 0644)
    if err != nil {
        fmt.Println("Error creating branch:", err)
        return
    }

    fmt.Printf("Created branch '%s'\n", name)
}

func ListBranches() {
    repoDir := ".vcs"
    refsDir := filepath.Join(repoDir, "refs")

    // Lire le répertoire des références
    files, err := ioutil.ReadDir(refsDir)
    if err != nil {
        fmt.Println("Error reading refs directory:", err)
        return
    }

    // Afficher les branches trouvées
    if len(files) == 0 {
        fmt.Println("No branches found.")
        return
    }

    fmt.Println("Branches:")
    for _, file := range files {
        fmt.Println(file.Name())
    }
}

func DeleteBranch(name string) {
    repoDir := ".vcs"
    refsDir := filepath.Join(repoDir, "refs")

    // Vérifier si la branche existe
    branchPath := filepath.Join(refsDir, name)
    if _, err := os.Stat(branchPath); os.IsNotExist(err) {
        fmt.Printf("Error: Branch '%s' does not exist.\n", name)
        return
    }

    // Supprimer le fichier de la branche
    err := os.Remove(branchPath)
    if err != nil {
        fmt.Println("Error deleting branch:", err)
        return
    }

    fmt.Printf("Deleted branch '%s'\n", name)
}
