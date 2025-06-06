package storage

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"myvcs/internal/objects"
)

type Repository struct {
    Path    string
    Head    string
    Stage   map[string]string
    Objects map[string]*objects.Commit
}

func InitRepo() {
	repoDir := ".vcs"
	subDirs := []string{
		"objects",          // Pour stocker les objets git
		"refs",             // Répertoire pour stocker les références de branches
		"refs/heads",       // Sous-répertoire pour stocker les branches spécifiques
	}


	if _, err := os.Stat(repoDir); !os.IsNotExist(err) {
		fmt.Println("Repository already initialized.")
		return
	}


	if err := os.Mkdir(repoDir, 0755); err != nil {
		fmt.Println("Error initializing repository:", err)
		return
	}


	for _, subDir := range subDirs {
		dirPath := filepath.Join(repoDir, subDir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			fmt.Println("Error creating subdirectory:", err)
			return
		}
	}


	headFile := filepath.Join(repoDir, "HEAD")
	defaultBranch := "refs/heads/main"
	if err := os.WriteFile(headFile, []byte(defaultBranch), 0644); err != nil {
		fmt.Println("Error creating HEAD file:", err)
		return
	}


	mainBranchFile := filepath.Join(repoDir, "refs/heads/main")
	if err := os.WriteFile(mainBranchFile, []byte(""), 0644); err != nil {
		fmt.Println("Error creating main branch file:", err)
		return
	}

	fmt.Println("Repository initialized.")
}

func (r *Repository) StageFile(path string) error {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }

    h := sha1.New()
    h.Write(content)
    hash := hex.EncodeToString(h.Sum(nil))

    r.Stage[filepath.Base(path)] = hash

    return ioutil.WriteFile(
        filepath.Join(r.Path, ".vcs/objects", hash),
        content,
        0644,
    )
}

func (r *Repository) Commit(message string) (*objects.Commit, error) {
    if len(r.Stage) == 0 {
        return nil, fmt.Errorf("nothing to commit")
    }

    commit := objects.NewCommit(message, r.Head, r.Stage)

    r.Objects[commit.Hash] = commit
    r.Head = commit.Hash

    headFile := filepath.Join(r.Path, ".vcs/HEAD")
    err := ioutil.WriteFile(headFile, []byte("refs/heads/main"), 0644)
    if err != nil {
        return nil, err
    }

    r.Stage = make(map[string]string)

    return commit, nil
}


func (r *Repository) CreateBranch(name string) error {
    branchFile := filepath.Join(r.Path, ".vcs/refs/heads", name)
    if _, err := os.Stat(branchFile); err == nil {
        return fmt.Errorf("branch '%s' already exists", name)
    }

    err := ioutil.WriteFile(branchFile, []byte(""), 0644)
    if err != nil {
        return fmt.Errorf("error creating branch: %w", err)
    }

    headFile := filepath.Join(r.Path, ".vcs/HEAD")
    err = ioutil.WriteFile(headFile, []byte("refs/heads/"+name), 0644)
    if err != nil {
        return fmt.Errorf("error updating HEAD: %w", err)
    }

    fmt.Printf("Branch '%s' created successfully\n", name)
    return nil
}
