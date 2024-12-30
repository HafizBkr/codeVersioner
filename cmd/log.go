package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"myvcs/internal/objects"
)

// Fonction pour afficher les commits
func ShowCommits() {
	repoDir := ".vcs"
	commitsDir := filepath.Join(repoDir, "commits")

	// Vérifier si le dépôt existe
	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		fmt.Println("Error: Repository not initialized. Run 'vcs start' first.")
		return
	}

	// Vérifier si le répertoire des commits existe
	if _, err := os.Stat(commitsDir); os.IsNotExist(err) {
		fmt.Println("No commits found.")
		return
	}

	// Lire le répertoire des commits
	files, err := os.ReadDir(commitsDir)
	if err != nil {
		fmt.Println("Error reading commits directory:", err)
		return
	}

	// Si aucun commit n'est trouvé
	if len(files) == 0 {
		fmt.Println("No commits found.")
		return
	}

	// Parcourir chaque fichier dans le répertoire des commits
	for _, file := range files {
		// Lire chaque fichier de commit (supposé être des fichiers JSON)
		commitPath := filepath.Join(commitsDir, file.Name())
		commitData, err := os.ReadFile(commitPath)
		if err != nil {
			fmt.Printf("Error reading commit file '%s': %v\n", file.Name(), err)
			continue
		}

		// Désérialiser les données en structure Commit
		var commit objects.Commit
		if err := json.Unmarshal(commitData, &commit); err != nil {
			fmt.Printf("Error unmarshalling commit file '%s': %v\n", file.Name(), err)
			continue
		}

		// Afficher les informations du commit (hash, message, date)
		commitDate := commit.Timestamp.Format(time.RFC1123) // Conversion de time.Time en une date lisible
		fmt.Printf("Commit: %s\nMessage: %s\nDate: %s\n\n", commit.Hash, commit.Message, commitDate)
	}
}
