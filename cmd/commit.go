package cmd

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"

    "myvcs/internal/objects"
)

func SaveCommit(message string) {
    repoDir := ".vcs"
    commitsDir := filepath.Join(repoDir, "commits")

    // Vérifier si le répertoire .vcs existe
    if _, err := os.Stat(repoDir); os.IsNotExist(err) {
        fmt.Println("Error: Repository not initialized. Run 'vcs start' first.")
        return
    }

    // Créer le répertoire des commits s'il n'existe pas
    if _, err := os.Stat(commitsDir); os.IsNotExist(err) {
        err = os.MkdirAll(commitsDir, 0755)
        if err != nil {
            fmt.Println("Error creating commits directory:", err)
            return
        }
    }

    // Charger le dernier commit depuis HEAD pour obtenir le parent
    headPath := filepath.Join(repoDir, "HEAD")
    var parentHash string
    if _, err := os.Stat(headPath); err == nil {
        parentBytes, err := os.ReadFile(headPath)
        if err == nil {
            parentHash = string(parentBytes)
        }
    }

    // Créer un nouvel objet Commit
    newCommit := objects.NewCommit(message, parentHash, nil)

    // Sérialiser le commit en JSON
    commitJSON, err := json.MarshalIndent(newCommit, "", "  ")
    if err != nil {
        fmt.Println("Error marshalling commit to JSON:", err)
        return
    }

    // Sauvegarder le commit dans un fichier
    commitPath := filepath.Join(commitsDir, newCommit.Hash+".json")
    err = os.WriteFile(commitPath, commitJSON, 0644)
    if err != nil {
        fmt.Println("Error saving commit file:", err)
        return
    }

    // Mettre à jour HEAD avec le nouveau hash du commit
    err = os.WriteFile(headPath, []byte(newCommit.Hash), 0644)
    if err != nil {
        fmt.Println("Error updating HEAD:", err)
        return
    }

    fmt.Printf("Commit saved: %s\n", newCommit.Hash)
}
