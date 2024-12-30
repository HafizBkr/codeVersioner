package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepo() {
	repoDir := ".vcs"
	// Liste des sous-répertoires nécessaires pour le dépôt
	subDirs := []string{
		"objects",         // Pour stocker les objets git
		"refs",            // Répertoire pour stocker les références de branches
		"refs/heads",      // Sous-répertoire pour stocker les branches spécifiques
		"staging",         // Répertoire pour stocker les fichiers en attente de commit
	}

	// Vérifier si le dépôt est déjà initialisé
	if _, err := os.Stat(repoDir); !os.IsNotExist(err) {
		fmt.Println("Repository already initialized.")
		return
	}

	// Créer le répertoire principal du dépôt .vcs
	if err := os.Mkdir(repoDir, 0755); err != nil {
		fmt.Println("Error initializing repository:", err)
		return
	}

	// Créer les sous-répertoires nécessaires
	for _, subDir := range subDirs {
		dirPath := filepath.Join(repoDir, subDir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			fmt.Println("Error creating subdirectory:", err)
			return
		}
	}

	// Créer le fichier HEAD par défaut pointant vers la branche "main"
	headFile := filepath.Join(repoDir, "HEAD")
	defaultBranch := "refs/heads/main"
	if err := os.WriteFile(headFile, []byte(defaultBranch), 0644); err != nil {
		fmt.Println("Error creating HEAD file:", err)
		return
	}

	// Créer la branche "main" dans le répertoire refs/heads
	mainBranchFile := filepath.Join(repoDir, "refs", "heads", "main")
	if err := os.WriteFile(mainBranchFile, []byte(""), 0644); err != nil {
		fmt.Println("Error creating main branch file:", err)
		return
	}

	fmt.Println("Repository initialized.")
}
